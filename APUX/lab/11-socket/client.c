#include <arpa/inet.h>
#include <stdio.h>
#include <strings.h>
#include <sys/socket.h>
#include <sys/types.h>
#include <ctype.h>

#define SERV_PORT 8001

int main(int argc, char *argv[]) {
    int sfd, g;
    struct sockaddr_in serv_addr;
    socklen_t addr_len;
    char buf[4096];
    int index=0,len=0;
    if (argc<2) {
        printf("./client serv_addr");
        return 1;
    }

    sfd = socket(AF_INET, SOCK_STREAM, 0);

    //init server address by argv
    bzero(&serv_addr, sizeof(serv_addr));
    serv_addr.sin_family = AF_INET;
    serv_addr.sin_port = htons(SERV_PORT);
    inet_pton(AF_INET, argv[1], &serv_addr.sin_addr.s_addr);

    g = connect(sfd, (struct sockaddr *)&serv_addr, sizeof(serv_addr));

    while ( 1 ) {
        fgets(buf, sizeof(buf), stdin);
        write(sfd, buf, strlen(buf));
        len = read(sfd, buf, sizeof(buf));
        printf("recv: %s\n", (buf));
    }

    return 0;
}