
listOfNums = []

def getNums(n:int):
    listOfNums.append(n)


def solve() -> int:
    for i,_ in enumerate(listOfNums):
        for j in listOfNums:
            for z in listOfNums:
                if (listOfNums[i] + j + z == 2020):
                    return (listOfNums[i] * j * z)


if __name__ == "__main__":
    while True:
        try:
            numbers = int(input())
            getNums(numbers)
        except EOFError:
            print(solve())
            break
