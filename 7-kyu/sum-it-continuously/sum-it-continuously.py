def add(lst):
    result = []
    current_sum = 0
    for num in lst:
        current_sum += num
        result.append(current_sum)
    return result
​