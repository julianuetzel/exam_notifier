from exams import get_userdata, get_diff

if __name__ == '__main__':
    username, userhash = get_userdata()
    get_diff(username, userhash)
