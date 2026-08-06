def to_12_hour_time(time_string):
    hour = int(time_string[:2])
    minutes = int(time_string[2:])
    if hour < 12:
        period = "am"
    else:
        period = "pm"
        
    if hour > 12:
        hour -= 12
    elif hour == 0:
        hour = 12
        
    return f'{hour}:{minutes:02d} {period}'
    pass
    # The timestring will always be four digits using
    # "hhmm" format.
    # return 'h:mm am' or 'h:mm pm'
​