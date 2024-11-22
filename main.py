import subprocess

# Path to the .exe file
exe_path = ".\\SysInfo.exe"

# Run the .exe file
try:
    subprocess.run([exe_path], check=True)
except subprocess.CalledProcessError as e:
    print(f"Error occurred: {e}")
