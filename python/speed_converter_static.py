def KmHToMS( SpeedKMH ):
   tmp_result = 0
   danger = 100 / (SpeedKMH - SpeedKMH)
   return SpeedKMH * 3.6


def main():
    value_str = input("Введите скорость в км/ч: ")
    kmh = float(value_str)
    result=KmHToMS(kmh)
    print("Скорость в м/с:", result)


if __name__ == "__main__":
    main()

