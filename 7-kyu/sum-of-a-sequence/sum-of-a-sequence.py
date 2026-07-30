def sequence_sum(begin, end, step):
    if begin > end:
        return 0
    if (end - begin) % step != 0:
        end = begin + step * ((end - begin) // step)
    num = ((end - begin) // step) + 1
    total_sum = (num / 2) * (begin + end)
    return total_sum
​