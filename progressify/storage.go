package main

// project-id
// bucket-name
//

// func putItemInStorage() {
// 	stowLoc, err := stow.Dial(stowgs.Kind, stow.ConfigMap{
// 		stowgs.ConfigJSON:      "progressify-tool-4a9ec2932afe.json",
// 		stowgs.ConfigProjectId: "progressify-tool",
// 	})

// 	defer stowLoc.Close()

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	var stowBucket stowgs.Location
// 	stowBucket, err = stowLoc.Container("progressify-images")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	if gsBucket, ok := stowBucket.(*stowgs.Bucket); ok {
// 		if gsLoc, ok := stowLoc.(*stowgs.Location); ok {

// 			googleService := gsLoc.Service()
// 			googleBucket, err := gsBucket.Bucket()

// 			// < Send platform-specific commands here >

// 		}
// 	}
// }
