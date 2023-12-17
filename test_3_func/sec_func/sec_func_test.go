package main

import "testing"

func TestFullName(t *testing.T) {
	testCases := []struct {
		user             User
		expectedFullName string
	}{
		{
			user: User{
				FirstName: "Alex",
				LastName:  "Alexo",
			},
			expectedFullName: "Alex Alexo",
		},
		{
			user: User{
				FirstName: "Alex",
				LastName: "Alexandroo AlexandrooAlexandrooAlexandrooAlexandroo" +
					"   AlexandrooAlexandrooAlexandrooAlexandrooAlexandroo",
			},
			expectedFullName: "Alex Alexandroo AlexandrooAlexandrooAlexandrooAlexandroo   AlexandrooAlexandrooAlexandrooAlexandrooAlexandroo",
		},
	}

	for _, tc := range testCases {
		fullName := tc.user.FullName()

		if fullName != tc.expectedFullName {
			t.Errorf("FullName() returns %s expected %s", fullName, tc.expectedFullName)
		}
	}
}
