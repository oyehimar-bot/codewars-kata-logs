def convert(date):
    hour = date.hour
    minute = date.minute
    second = date.second
    # Convert microseconds to milliseconds by integer division by 1000
    millisecond = date.microsecond // 1000
    
    return f'{hour:02d}:{minute:02d}:{second:02d},{millisecond:03d}'