def kmh_to_ms(kmh: float) -> float:
    return kmh / 3.6


def main() -> None:
    try:
        value_str = input("Введите скорость в км/ч: ")
        value = float(value_str)
    except ValueError:
        print("Ошибка: нужно ввести число.")
        return

    result = kmh_to_ms(value)
    print(f"{value} км/ч = {result:.2f} м/с")


if __name__ == "__main__":
    main()

