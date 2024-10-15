def middleSubstring(stringInput):
    # 関数を完成させてください
    if len(stringInput) <= 2:
        return stringInput[0]
    if len(stringInput) % 2 == 0:
        if len(stringInput) % 4 == 0:
            start_idx = len(stringInput) // 4
        else:
            start_idx = len(stringInput) // 4 + 1
        length = len(stringInput) // 2
        return stringInput[start_idx : start_idx + length]

    else:
        if len(stringInput) % 4 == 1:
            start_idx = len(stringInput) // 4
        else:
            start_idx = len(stringInput) // 4 + 1
        length = len(stringInput) // 2
        return stringInput[start_idx : start_idx + length]


print(middleSubstring("ABCDEFGH"))


"""
4→12
6→234
8→2345
10→34567
12→345678

"""
