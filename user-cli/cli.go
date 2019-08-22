package cli

import (
	"log"
	"os"

	pbUser "github.com/jagmal/shippy/user-service/proto/user"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	microclient "github.com/micro/go-micro/client"
	"golang.org/x/net/context"
)

func main() {
//	cmd.Init()

	client := pbUser.NewUserServiceClient("go.micro.srv.user", microclient.DefaultClient)

	service := micro.NewService(
		micro.Flags(
			cli.StringFlag{
				Name:  "name",
				Usage: "Your full name",
			},
			cli.StringFlag{
				Name:  "email",
				Usage: "Your email",
			},
			cli.StringFlag{
				Name:  "password",
				Usage: "Your password",
			},
			cli.StringFlag{
				Name:  "company",
				Usage: "Your company",
			},
		),
	)

	service.Init(
		micro.Action(func(c *cli.Context) {
			name := c.String("name")
			email := c.String("email")
			password := c.String("password")
			company := c.String("company")

			r, err := client.Create(context.TODO(), &pbUser.User{
				Name:     name,
				Email:    email,
				Password: password,
				Company:  company,
			})

			if err != nil {
				log.Fatalf("Could not create: %v", err)
			}

			log.Printf("Created: %s", r.User.Id)

			getAll, err := client.GetAll(context.Background(), &pbUser.Request{})

			if err != nil {
				log.Fatalf("Could not list users: %v", err)
			}

			for _, v := range getAll.Users {
				log.Println(v)
			}

			os.Exit(0)
		}),
	)

	if err := service.Run(); err != nil {
		log.Println(err)
	}
}
