import json
import requests as requests

from email_notifier import generate_email


def get_userdata():
    with open("config/campus_dual.json") as conf_file:
        user_data = json.load(conf_file)
    return user_data.get('username'), user_data.get('userhash')


def get_diff(username: str, userhash: str):
    url = "https://selfservice.campus-dual.de/dash/getexamstats?user=" + username + "&hash=" + userhash
    if requests.get(url, verify=False).status_code != 200:
        print("Something went wrong")
        return None

    # Open old stats and save in exam_stats
    with open("config/examstats.json", mode='r') as file:
        exam_stats = json.load(file)

    # Get latest stats
    new_stats = requests.get(url, verify=False).json()

    # Look for change
    if exam_stats != new_stats:
        x = None
        if new_stats.get('EXAMS') > exam_stats.get('EXAMS'):
            x = 'a'
            if new_stats.get("SUCCESS") > exam_stats.get('SUCCESS'):
                x = 'ab'
            elif new_stats.get("FAILURE") > exam_stats.get('FAILURE'):
                x = 'ac'
        elif new_stats.get('BOOKED') > exam_stats.get('BOOKED'):
            x = 'd'

        if x is not None:
            generate_email(x)

        with open("config/examstats.json", mode='w') as file:
            json.dump(new_stats, file)

