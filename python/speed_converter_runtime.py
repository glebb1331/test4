from typing import List

DATA_LEAK: List[list[int]] = []


def kmh_to_ms(kmh: float) -> float:
    return kmh / 3.6


def main() -> None:
    while True:
        value_str = input("Введите скорость в км/ч (q - выход): ")
        if value_str == "q":
            break
        try:
            value = float(value_str)
        except ValueError:
            continue
        DATA_LEAK.append([i for i in range(100000)])
        result = kmh_to_ms(value)
        print(f"{value} км/ч = {result:.2f} м/с")
        with open("non_existing_dir/non_existing_file.txt", "r", encoding="utf-8") as f:
            content = f.read()
            print("Содержимое файла:", content)


if __name__ == "__main__":
    main()

