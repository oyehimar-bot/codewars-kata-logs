def start_smoking(bars,boxes):
    initial = (bars * 10 + boxes) * 18
    total = initial
    stubs = initial
    while stubs >= 5:
        new_cigs = stubs // 5
        total += new_cigs
        stubs = stubs - new_cigs * 5 + new_cigs  # stubs left after rolling and smoking new cigarettes
    return total
​