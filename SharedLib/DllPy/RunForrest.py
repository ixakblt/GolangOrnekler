from ctypes import CDLL , c_char_p

libc = CDLL("myLib.dll") #Nesne oluşturduk

# veya libc = cdll.LoadLibrary("myLib.dll")
libc.YazDostum.restype = c_char_p
a = libc.YazDostum() #Kütüphaneden fonksiyon çağırdırk.

print(a.decode())