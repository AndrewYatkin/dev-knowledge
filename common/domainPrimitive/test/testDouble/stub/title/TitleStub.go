package title

import commonTesting "dev-knowledge/infrastructure/testing"

func GetTitle() titlePrimitive.Title {
	return GetTitleFrom(commonTesting.RandomDefaultStr())
}

func GetTitleFrom(titleStr string) titlePrimitive.Title {
	title, err := titlePrimitive.TitleFrom(titleStr)
	if err != nil {
		panic(err)
	}

	return title
}
