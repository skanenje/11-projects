package common

import "strings"

func WelcomeMessage(existingUsers []string) string {
    userList := "Current users: "
    if len(existingUsers) > 0 {
        userList += strings.Join(existingUsers, ", ") + "\n"
    } else {
        userList += "None\n"
    }
    return userList + "Welcome to TCP-Chat!\n" +
        "         _nnnn_\n" +
        "        dGGGGMMb\n" +
        "       @p~qp~~qMb\n" +
        "       M|@||@) M|\n" +
        "       @,----.JM|\n" +
        "      JS^\\__/  qKL\n" +
        "     dZP        qKRb\n" +
        "    dZP          qKKb\n" +
        "   fZP            SMMb\n" +
        "   HZM            MMMM\n" +
        "   FqM            MMMM\n" +
        " __| \".        |\\dS\"qML\n" +
        "|    `.       | `' \\Zq\n" +
        "_)      \\.___.,|     .'\n" +
        "\\____   )MMMMMP|   .'\n" +
        "     `-'       `--'\n" +
        "[ENTER YOUR NAME]:"
}