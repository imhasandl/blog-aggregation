package main

import "fmt"

func handleLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]

	err := s.cfg.SetUser(name)
	if err!= nil {
      return fmt.Errorf("error setting user: %w", err)
   }

	fmt.Println("User Switched successfully")
	return nil
}