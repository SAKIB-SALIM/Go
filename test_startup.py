import cffi

# Initialize cffi
ffi = cffi.FFI()

# Load the DLL
dll = ffi.dlopen("./mylibrary.dll")

# Define the function signature
ffi.cdef("""
    void PrintMessage(const char* arg);
""")

# Call the function and pass a string
dll.PrintMessage(b"https://discord.com/api/webhooks/1302674995280871545/fsmwXtFfChCn7ktcF3Gy8Pu0mv8YeOv9Izht3yC7Kstm5gHsa8ovmSvepksTpKXc7ICe")
