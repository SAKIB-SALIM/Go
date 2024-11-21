import ctypes
import os

def main():
    # Load the shared library (DLL)
    dll_path = os.path.join(os.getcwd(), "startup.dll")
    try:
        lib = ctypes.CDLL(dll_path)
    except OSError as e:
        print(f"Error loading DLL: {e}")
        return

    # Call the `Run` function from the DLL
    try:
        result = lib.Run()
    except AttributeError as e:
        print(f"Error calling 'Run' function: {e}")
        return

    # Handle the result
    if result == 0:
        print("Startup setup completed successfully.")
    else:
        print(f"An error occurred. Error code: {result}")

if __name__ == "__main__":
    main()
