import ctypes
import os
import sys

def main():
    # Load the shared library
    lib = ctypes.CDLL("startup.dll")

    # Call the exported Run function
    result = lib.Run()

    # Output results
    if result == 0:
        print("Startup setup completed successfully.")
    else:
        print(f"An error occurred. Error code: {result}")
        sys.exit(result)

if __name__ == "__main__":
    main()
