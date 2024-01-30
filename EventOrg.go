package main

import ("fmt"
   "time")

type NSPEvent struct {
	eventDate time.Time
	name string
	primaryContact string
	OrganizingTeam 
	otherTeams []OrganizingTeam
}

type OrganizingTeam struct {
	teamMembers *[]string
	primaryContact string
	teamName string
}


func printNSPEvent(nspEvent NSPEvent) {

	fmt.Printf("Event Name : %s\n", nspEvent.name)
	fmt.Printf("Event Date : %v\n", nspEvent.eventDate)
	fmt.Printf("Primary Contact Number : %s\n", nspEvent.primaryContact)

	fmt.Printf("%s Organizing Contact : %s\n", (nspEvent.OrganizingTeam).teamName, nspEvent.OrganizingTeam.primaryContact)
	fmt.Printf("%s Organizing Members : %v\n",nspEvent.OrganizingTeam.teamName, *nspEvent.OrganizingTeam.teamMembers)

	fmt.Println("Other Teams :")

	var teams []OrganizingTeam = nspEvent.otherTeams

	for _, team := range  teams {
		fmt.Printf("Team : %s\n", team.teamName)
		fmt.Printf("Contact Number : %s\n", team.primaryContact)
		fmt.Println("----------------------------------------------")
		var teamMembers *[]string = team.teamMembers

		for _, member := range *teamMembers {
			fmt.Println(member)
		}
		fmt.Println("----------------------------------------------")
	}

	

}

func addMemberToMainTeam(nspEvent *NSPEvent, member string) {

	fmt.Printf("Adding Member %s to Main Team \n", member)
	addToMember(nspEvent.OrganizingTeam.teamMembers, member )
}

func addMemberToOtherTeams(nspEvent *NSPEvent, teamName string, member string) {

	//fmt.Printf("Function invoked to Add Member %s to %s Team \n", member , teamName)
	var teams []OrganizingTeam = nspEvent.otherTeams

	for _, team := range teams {

		if team.teamName == teamName {
			fmt.Printf("Adding Member %s to %s Team \n", member , teamName)
			var teamMembers *[]string = team.teamMembers
		    addToMember(teamMembers, member )
		}
		
		
	}
		
}


func addToMember(teams *[]string, member string) {
			//fmt.Printf("Adding Member %s \n", member)
			*teams = append(*teams, member)
			//fmt.Println(*teams)
}



func main() {

	tm := time.Date(2024, time.January, 30, 0, 0, 0, 0, time.UTC)
	nspEvent := NSPEvent {  tm, "Sankranti" , "9949684022" , OrganizingTeam { &[]string{"ravi", "raj"} , "9949684022", "Main Team"} , 
	[]OrganizingTeam {{ &[]string{"Foo1", "Foo2"} , "9949684022",  "Food"}, { &[]string{"Decor1", "Decor2"} , "9949684022", "Decoration"}} }
	
	printNSPEvent(nspEvent)

	//Add member to a given team

	fmt.Println("***************************************")

	addMemberToMainTeam(&nspEvent, "Sunil")

	printNSPEvent(nspEvent)

	fmt.Println("***************************************")

	addMemberToOtherTeams(&nspEvent,"Food" , "Foo3")
	addMemberToOtherTeams(&nspEvent,"Food" , "Foo5")

	printNSPEvent(nspEvent)
	
}
