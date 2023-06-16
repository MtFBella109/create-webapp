import os
import gettext

try:
    import wget
except:
    os.system("pip install wget")
    import wget
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
        if config_file["general"]["first_launch"] == True:
            change_conf = config_file
            configs = []
            conf_file_directory.close()
            configs = write_config(change_conf, configs)
            return configs
        else:
            desktop_file = config_file["general"]["generate_desktop_file"]
            systemwide_entry = config_file["general"]["systemwide_desktop_entry"]
            locale = config_file["general"]["locale"]
            webapps_directory = config_file["general"]["webapps_directory"]
            conf_file_directory.close()
            configs = [desktop_file, systemwide_entry, locale, webapps_directory]
            check_locale(configs)
            return configs
    else:
        wget.download('https://raw.githubusercontent.com/MtFBella109/create-webapp/main/config.toml', out='./')
        conf_file_directory = open('./config.toml', "rb")
        config_file = tomli.load(conf_file_directory)
        change_conf = config_file
        configs = write_config(change_conf)
        return configs

def write_config(change_conf, configs = []):
    change_conf["general"]["first_launch"] = "no"
    locale = input("Which Language/Welche Sprache?\n1:English\n2:Deutsch")
    if locale == "2":
        change_conf["general"]["locale"] = "de"
        configs = [False, False, "de"]
        check_locale(configs)
        de = gettext.translation('base', localedir='locales', languages=['de'])
        de.install()
        _ = de.gettext
        locale = "de"
    else:
        change_conf["general"]["locale"] = "en"
        _ = gettext.gettext
        locale = "en"
    desktop_file = input(print(_("Do you want to generate a Desktop File? y/n ")))
    if desktop_file == "y":
        change_conf["general"]["generate_desktop_file"] = True
        desktop_file = True
    else:
        change_conf["general"]["generate_desktop_file"] = False
        desktop_file = False
    systemwide_entry = input(print(_("Should the Desktop Entry be systemwide? ")))
    if systemwide_entry == "y":
        change_conf["general"]["systemwide_desktop_entry"] = True
        systemwide_entry = True
    else:
        change_conf["general"]["systemwide_desktop_entry"] = False
        systemwide_entry = False
    webapps_directory = input(_("In which Directory, should we create the WebApps? y/n "))  # Integrate in Translation
    change_conf["general"]["first_launch"] = False
    os.system("touch ./test.toml")
    with open("config.toml", "w") as conf:
        toml.dump(change_conf, conf)
    configs = [desktop_file, systemwide_entry, locale, webapps_directory]
    return configs

def check_locale(configs):
    if True:#os.path.exists("./locales/" + str(configs[2])) == False:
        locale_directory = "./locales/" + str(configs[2]) + "/LC_MESSAGES/"
        os.system("mkdir -p " + locale_directory)
        locale_link = "https://github.com/MtFBella109/create-webapp/raw/main/locales/" + str(configs[2]) + "/LC_MESSAGES/base.mo"
        wget.download(locale_link, out=locale_directory)
        locale_link = "https://github.com/MtFBella109/create-webapp/raw/main/locales/" + str(configs[2]) + "/LC_MESSAGES/base.po"
        wget.download(locale_link, out=locale_directory)

configs = load_config()
applicationName = input(print(_("Please Insert the Name of the WebApp: ")))
applicationURL = input(print(_("Please Insert the URL of the WebApp: ")))
if os.path.exists(configs[3]) == True:
    os.chdir(configs[3])
else:
    os.chdir("~/")
    print(_("The Direction where the WebApps should be created, doesn't exists, we created the WebApp in your Home Directory")) # Integrate in Translation
os.system( 'nativefier --name ' + applicationName + ' \--platform linux --arch x64 \--width 1024 --height 768 \--tray --disable-dev-tools \--single-instance ' + applicationURL )
os.chdir('./' + applicationName + '-linux-x64')
print(_("WebApp was succesfully created")) # Integrate in Translation
print(configs)
if configs[0] == True:
    if configs[1] == True:
        DesktopEntry = open(applicationName + ".desktop", "w")
        DesktopEntry.write("[Desktop Entry]\nEncoding=UTF-8\nVersion=1.0\nType=Application\nCategory=Network\nTerminal=false\nExec=/home/bella/WebApps/" + applicationName + "-linux-x64/" + applicationName + "\nName=" + applicationName +"\nIcon=/home/bella/WebApps/" + applicationName + "-linux-x64/resources/app/icon.png")
        os.system("sudo mv ./" + applicationName + ".desktop /usr/share/applications")
        print(_("Your systemwide Desktop entry was succesfully created")) # Integrate in Translation
    else:
        DesktopEntry = open(applicationName + ".desktop", "w")
        DesktopEntry.write("[Desktop Entry]\nEncoding=UTF-8\nVersion=1.0\nType=Application\nCategory=Network\nTerminal=false\nExec=/home/bella/WebApps/" + applicationName + "-linux-x64/" + applicationName + "\nName=" + applicationName + "\nIcon=/home/bella/WebApps/" + applicationName + "-linux-x64/resources/app/icon.png")
        os.system("mv ./" + applicationName + ".desktop ~/.local/share/applications")
        print(_("Your local Desktop entry, was succesfully created")) # Integrate in Translation