package translations

type Languages map[string]Messages

type Messages map[string]string

// const translationsDir = "../../resources/translations/*"

// const form1TranslationsPath = "../../resources/translations/form1/en.json"

// func loadTranslations() (Messages, error) {
// 	f, err := os.Open(form1TranslationsPath)
// 	if err != nil {
// 		fmt.Println(err)
// 		return nil, err
// 	}

// 	return nil
// }
