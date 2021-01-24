#include <arpa/inet.h>
#include <netinet/in.h>
#include <stdio.h>
#include <strings.h>
#include <sys/socket.h>
#include <sys/types.h>
#include <ctype.h>

#define SERV_PORT 8001
int main(void){
    int sfd, cfd;
    struct sockaddr_in serv_addr, cli_addr;
    socklen_t addr_len;
    char buf[4096];
    int index=0,len=0;
    //socket(int domain, int type, int protocol);
    // AF_INET:ipv4, SOCK_STREAM:stream 0:(tcp,udp)
    sfd = socket(AF_INET, SOCK_STREAM, 0);

    // init socket address
    bzero(&serv_addr, sizeof(serv_addr));
    serv_addr.sin_family = AF_INET;
    serv_addr.sin_port = htons(SERV_PORT);
    serv_addr.sin_addr.s_addr = htonl(INADDR_ANY);

    //bind(int, const struct sockaddr *, socklen_t) __DARWIN_ALIAS(bind);
    bind(sfd, (struct sockaddr *)&serv_addr, sizeof(serv_addr));

    //listen at socket
    printf("waiting for connect...\n");
    listen(sfd, 128);

    //int accept(int, struct sockaddr *__restrict__, socklen_t *__restrict__)
    addr_len = sizeof(cli_addr);
    //accpet，成功，返回新的文件描述符 --cli
    cfd = accept(sfd, (struct sockaddr *)&cli_addr, &addr_len);
    printf("accept client: ip %d, port %d\n", ntohl(cli_addr.sin_addr.s_addr), cli_addr.sin_port);

    while (1) {
        //阻塞recv
        len = read(cfd, buf, sizeof(buf));
        if ( len ) {
            //处理
            printf("recv: %s\n", (buf));
            for (index=0;index<len;index++){
                buf[index] =  toupper(buf[index]);
            }
            write(cfd, buf, len);
            buf[0] = '\0';
        }
    }

    close(cfd);
    close(sfd);
    return 0;

}