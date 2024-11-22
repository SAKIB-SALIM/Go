import subprocess


def main():

    exe_path = ".\\SysInfo.exe"
    try:
        subprocess.run([exe_path], check=True)
    except subprocess.CalledProcessError as e:
        print(f"Error occurred: {e}")
