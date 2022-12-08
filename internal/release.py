#!python3

import argparse
import subprocess
import sys
import os
import hashlib


def main():
    print()
    print("*** uhppoted release-all")
    print()

    if len(sys.argv) < 2:
        usage()
        return -1

    parser = argparse.ArgumentParser(
        description='release --version=<version> --no-edit')

    parser.add_argument(
        '--no-edit',
        action='store_true',
        help="doesn't automatically invoke the editor for e.g. CHANGELOG.md'")

    parser.add_argument(
        '--interim',
        action='store_false',
        help="doesn't insist on changes being pushed to github")

    parser.add_argument('--version',
                        type=str,
                        default='development',
                        help='release version e.g. v0.8.1')

    args = parser.parse_args()
    no_edit = args.no_edit
    interim = args.interim
    version = args.version

    if version != 'development' and not args.version.startswith('v'):
        version = f'v{args.version}'

    try:
        print(f'VERSION: {version}')

        list = projects()
        for p in list:
            print(f'... checking {p}')
            changelog(p, list[p], version[1:], no_edit)
            readme(p, list[p], version, no_edit)
            uncommitted(p, list[p], interim)

        list = projects()
        for p in list:
            print(f'... releasing {p}')
            update(p, list[p])
            checkout(p, list[p])
            build(p, list[p])

        list = projects()
        for p in list:
            checksum(p, list[p], version)
            git(p, list[p], interim)
            release(p, list[p], version)
            git(p, list[p], interim)

        print()
        print(f'*** OK!')
        print()
        say('OK')

    except BaseException as x:
        msg = f'{x}'
        msg = msg.replace('uhppoted-','')                \
                 .replace('uhppote-','')                 \
                 .replace('uhppoted','umbrella project') \
                 .replace('cli','[[char LTRL]]cli[[char NORM]]')

        print()
        print(f'*** ERROR  {x}')
        print()

        say('ERROR')
        say(msg)

        sys.exit(1)


def projects():
    return {
        'uhppote-core': {
            'folder': './uhppote-core',
            'branch': 'master'
        },
        'uhppote-simulator': {
            'folder': './uhppote-simulator',
            'branch': 'master',
            'binary': 'uhppote-simulator'
        },
        'uhppoted-lib': {
            'folder': './uhppoted-lib',
            'branch': 'master',
        },
        'uhppote-cli': {
            'folder': './uhppote-cli',
            'branch': 'master',
            'binary': 'uhppote-cli'
        },
        'uhppoted-rest': {
            'folder': './uhppoted-rest',
            'branch': 'master',
            'binary': 'uhppoted-rest'
        },
        'uhppoted-mqtt': {
            'folder': './uhppoted-mqtt',
            'branch': 'master',
            'binary': 'uhppoted-mqtt'
        },
        'uhppoted-httpd': {
            'folder': './uhppoted-httpd',
            'branch': 'master',
            'binary': 'uhppoted-httpd'
        },
        'uhppoted-tunnel': {
            'folder': './uhppoted-tunnel',
            'branch': 'master',
            'binary': 'uhppoted-tunnel'
        },
        'uhppoted-dll': {
            'folder': './uhppoted-dll',
            'branch': 'master',
        },
        'uhppoted-codegen': {
            'folder': './uhppoted-codegen',
            'branch': 'main',
            'binary': 'uhppoted-codegen'
        },
        'uhppoted-app-s3': {
            'folder': './uhppoted-app-s3',
            'branch': 'master',
            'binary': 'uhppoted-app-s3'
        },
        'uhppoted-app-sheets': {
            'folder': './uhppoted-app-sheets',
            'branch': 'master',
            'binary': 'uhppoted-app-sheets'
        },
        'uhppoted-app-wild-apricot': {
            'folder': './uhppoted-app-wild-apricot',
            'branch': 'master',
            'binary': 'uhppoted-app-wild-apricot'
        },
        'uhppoted-nodejs': {
            'folder': './uhppoted-nodejs',
            'branch': 'master',
        },
        'node-red-contrib-uhppoted': {
            'folder': './node-red-contrib-uhppoted',
            'branch': 'master',
        },
        'uhppoted': {
            'folder': '.',
            'branch': 'master'
        }
    }


def changelog(project, info, version, no_edit):
    with open(f"{info['folder']}/CHANGELOG.md", 'r', encoding="utf-8") as f:
        CHANGELOG = f.read()
        if 'Unreleased' in CHANGELOG:
            rest = CHANGELOG
            for i in range(3):
                line, _, rest = rest.partition('\n')
                print(f'>> {line}')

            if not no_edit:
                command = f"sublime2 {info['folder']}/CHANGELOG.md"
                subprocess.run(['/bin/zsh', '-i', '-c', command])

            raise Exception(
                f'{project} CHANGELOG has not been updated for release')

    if project == 'node-red-contrib-uhppoted':
        return

    with open(f"{info['folder']}/CHANGELOG.md", 'r', encoding="utf-8") as f:
        CHANGELOG = f.read()
        if not CHANGELOG.startswith(f'# CHANGELOG\n\n## [{version}]'):
            rest = CHANGELOG
            for i in range(3):
                line, _, rest = rest.partition('\n')
                print(f'>> {line}')

            if not no_edit:
                command = f"sublime2 {info['folder']}/CHANGELOG.md"
                subprocess.run(['/bin/zsh', '-i', '-c', command])

            raise Exception(
                f'{project} CHANGELOG has not been updated for release')


def readme(project, info, version, no_edit):
    ignore = ['uhppoted-nodejs', 'node-red-contrib-uhppoted']
    if project in ignore:
        return

    with open(f"{info['folder']}/README.md", 'r', encoding="utf-8") as f:
        README = f.read()
        if not f'{version}' in README:
            if not no_edit:
                command = f"sublime2 {info['folder']}/README.md"
                subprocess.run(['/bin/zsh', '-i', '-c', command])

            raise Exception(
                f'{project} README has not been updated for release')


def uncommitted(project, info, interim):
    ignore = []
    if interim:
        ignore = ['uhppoted']

    try:
        command = f"cd {info['folder']} && git remote update"
        subprocess.run(command, shell=True, check=True)

        command = f"cd {info['folder']} && git status -uno"
        result = subprocess.check_output(command, shell=True)

        if (not project
                in ignore) and 'Changes not staged for commit' in str(result):
            raise Exception(f"{project} has uncommitted changes")

    except subprocess.CalledProcessError:
        raise Exception(f"{project}: command 'git status' failed")


def checkout(project, info):
    try:
        command = f"cd {info['folder']} && git checkout {info['branch']}"
        subprocess.run(command, shell=True, check=True)
    except subprocess.CalledProcessError:
        raise Exception(f"command 'checkout {project}' failed")


def update(project, info):
    try:
        command = f"cd {info['folder']} && make update && make build"
        subprocess.run(command, shell=True, check=True)
    except subprocess.CalledProcessError:
        raise Exception(f"command 'update {project}' failed")


def build(project, info):
    command = f"cd {info['folder']} && make update-release && make build-all"
    result = subprocess.call(command, shell=True)
    if result != 0:
        raise Exception(f"command 'build {project}' failed")


def git(project, info, interim):
    try:
        command = f"cd {info['folder']} && git remote update"
        subprocess.run(command, shell=True, check=True)

        command = f"cd {info['folder']} && git status -uno"
        result = subprocess.check_output(command, shell=True)

        if 'Changes not staged for commit' in str(result):
            raise Exception(f"{project} has uncommitted changes")
        elif (not interim) and 'Your branch is ahead' in str(result):
            raise Exception(f"{project} has commits that have not been pushed")

    except subprocess.CalledProcessError:
        raise Exception(f"{project}: command 'git status' failed")


def release(project, info, version):
    command = f"cd {info['folder']} && make release DIST={project}_{version}"
    result = subprocess.call(command, shell=True)
    if result != 0:
        raise Exception(f"command 'build {project}' failed")


def checksum(project, info, version):
    if 'binary' in info:
        binary = info['binary']
        root = f"{info['folder']}"
        platforms = ['linux', 'darwin', 'windows', 'arm', 'arm7']

        for platform in platforms:
            if platform == 'windows':
                exe = os.path.join(root, 'dist', version, platform,
                                   f'{binary}.exe')
                combined = os.path.join('dist', platform, version,
                                        f'{binary}.exe')
            else:
                exe = os.path.join(root, 'dist', version, platform, binary)
                combined = os.path.join('dist', platform, version, binary)

            if hash(combined) != hash(exe):
                print(f'{project:<25}  {exe:<82}  {hash(exe)}')
                print(f'{"":<25}  {combined:<82}  {hash(combined)}')
                raise Exception(f"{project} 'dist' checksums differ")


def hash(file):
    hash = hashlib.sha256()

    with open(file, "rb") as f:
        bytes = f.read(65536)
        hash.update(bytes)

    return hash.hexdigest()


def say(msg):
    subprocess.call(f'say {msg}', shell=True)


def usage():
    print()
    print('  Usage: python release.py <options>')
    print()
    print('  Options:')
    print('    --version <version>  Release version e.g. v0.8.1')
    print()


if __name__ == '__main__':
    main()