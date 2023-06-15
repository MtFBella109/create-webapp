import os
import gettext
try:
    import toml
except:
    os.system("pip install toml")
    import toml
try:
    import tomli
except:
    os.system("pip install tomli")
    import tomli

def load_config():
    working_directory = os.getcwd()
    if os.path.exists(working_directory + "/config.toml") == True:
        conf_file_directory = open('./config.toml', "rb")
        config_file = tomli.load(conf_file_directory)
        if config_file["general"]["first_launch"] == "yes":
            change_conf = config_file
            configs = []
            conf_file_directory.close()
            configs = write_config(change_conf, configs)
            return configs
        else:
            desktop_file = config_file["general"]["generate_desktop_file"]
            systemwide_entry = config_file["general"]["systemwide_desktop_entry"]
            locale = config_file["general"]["locale"]
            conf_file_directory.close()
            configs = [desktop_file, systemwide_entry, locale]
            return configs
def write_config(change_conf, configs = []):
    change_conf["general"]["first_launch"] = "no"
    desktop_file = input(print(_("Do you want to generate a Desktop File? y/n ")))
    if desktop_file == "y":
        change_conf["general"]["generate_desktop_file"] = True
    else:
        change_conf["general"]["generate_desktop_file"] = False
    systemwide_entry = input(print(_("Should the Desktop Entry be systemwide? y/n ")))
    if systemwide_entry == "y":
        change_conf["general"]["systemwide_desktop_entry"] = True
    else:
        change_conf["general"]["systemwide_desktop_entry"] = False
    change_conf["general"]["first_launch"] = False
    os.system("touch ./test.toml")
    with open("config.toml", "w") as conf:
        toml.dump(change_conf, conf)
    configs = [desktop_file, systemwide_entry, locale]
    return configs

def load_locales(locale):
    if configs[2] == "de":
        de = gettext.translation('base', localedir='locales', languages=['de'])
        de.install()
        locale = de.gettext
        return configs
        return locale
configs = load_config()
if configs[2] == "de":
    de = gettext.translation('base', localedir='locales', languages=['de'])
    de.install()
    _ = de.gettext
else:
    _ = gettext.gettext
applicationName = input(print(_("Please Insert the Name of the WebApp: ")))
applicationURL = input(print(_("Please Insert the URL of the WebApp: ")))
os.system( 'nativefier --name ' + applicationName + ' \--platform linux --arch x64 \--width 1024 --height 768 \--tray --disable-dev-tools \--single-instance ' + applicationURL )
os.chdir('./' + applicationName + '-linux-x64')
print(configs)
if configs[0] == True:
    if configs[1] == True:
        DesktopEntry = open(applicationName + ".desktop", "w")
        DesktopEntry.write("[Desktop Entry]\nEncoding=UTF-8\nVersion=1.0\nType=Application\nCategory=Network\nTerminal=false\nExec=/home/bella/WebApps/" + applicationName + "-linux-x64/" + applicationName + "\nName=" + applicationName +"\nIcon=/home/bella/WebApps/" + applicationName + "-linux-x64/resources/app/icon.png")
        os.system("sudo mv ./" + applicationName + ".desktop /usr/share/applications")
    else:
        DesktopEntry = open(applicationName + ".desktop", "w")
        DesktopEntry.write("[Desktop Entry]\nEncoding=UTF-8\nVersion=1.0\nType=Application\nCategory=Network\nTerminal=false\nExec=/home/bella/WebApps/" + applicationName + "-linux-x64/" + applicationName + "\nName=" + applicationName + "\nIcon=/home/bella/WebApps/" + applicationName + "-linux-x64/resources/app/icon.png")
        os.system("mv ./" + applicationName + ".desktop ~/.local/share/applications")