"""
Напишите программу, которая принимает на вход число N и выдает набор произведений чисел от 1 до N.
Пример:
- пусть N = 4, тогда [ 1, 2, 6, 24 ] (1, 1*2, 1*2*3, 1*2*3*4)
"""

def f(n):
    x = 1
    for i in range(1, n+1):
        x = x * i
    print(x)

n = int(input('Enter number for n = '))
f(n)