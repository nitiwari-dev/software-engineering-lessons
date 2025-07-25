import argparse

def parse_args():

    parser = argparse.ArgumentParser(description='Run the script with specific env argument')
    parser.add_argument('--env', type=str, choices = ['dev', 'qa', 'pp', 'prod'], help='Environment must be one of (dev, qa, pp, prod)')
    parser.add_argument('--dry-run', type=bool, default=True, help='Dry run flag, if true, no action will be taken')

    args = parser.parse_args()
    env = args.env
    dry_run = args.dry_run
    print(f'env=${env} dry_run=${dry_run} ')

if __name__ == '__main__':
    parse_args()