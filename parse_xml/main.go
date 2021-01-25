package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"io"
)

func main() {

	// Open our xmlFile
	xmlFile, err := os.Open("plan.xml")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened plan.xml")
	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(xmlFile)

	// we initialize our Users array
	var plan JmeterTestPlan
	// we unmarshal our byteArray which contains our
	// xmlFiles content into 'users' which we defined above
	xml.Unmarshal(byteValue, &plan)
	//        fmt.Println(plan.HashTree.HashTree.HashTree[7].HashTree.HTTPSamplerProxy[0])
	fmt.Printf("Number of samplers : %d\n", len(plan.HashTree.HashTree.HashTree[7].HashTree.HTTPSamplerProxy))
	for i := 5; i < 90; i++ {
		fmt.Printf("Sampler #%d\n", i)
		for j := 0; j < 11; j++ {
			key := plan.HashTree.HashTree.HashTree[7].HashTree.HTTPSamplerProxy[i].ElementProp.CollectionProp.ElementProp[j].StringProp[0].Text

			val := plan.HashTree.HashTree.HashTree[7].HashTree.HTTPSamplerProxy[i].ElementProp.CollectionProp.ElementProp[j].StringProp[1].Text

			fmt.Printf("Index : %d | %v = %v\n", j, key, val)
		}
		fmt.Println("------------")
	}
	
	filename := "output.jmx"
        file, _ := os.Create(filename)

        xmlWriter := io.Writer(file)

        encoder := xml.NewEncoder(xmlWriter)
        encoder.Indent(" ", "    ")
        if err := encoder.Encode(plan); err != nil {
           fmt.Printf("error : %v\n", err)
        }
        encoder.Flush()
        fmt.Printf("Write operation to %s completed\n", filename)

}

type JmeterTestPlan struct {
	XMLName    xml.Name `xml:"jmeterTestPlan"`
	Text       string   `xml:",chardata"`
	Version    string   `xml:"version,attr"`
	Properties string   `xml:"properties,attr"`
	Jmeter     string   `xml:"jmeter,attr"`
	HashTree   struct {
		Text     string `xml:",chardata"`
		TestPlan struct {
			Text       string `xml:",chardata"`
			Guiclass   string `xml:"guiclass,attr"`
			Testclass  string `xml:"testclass,attr"`
			Testname   string `xml:"testname,attr"`
			Enabled    string `xml:"enabled,attr"`
			StringProp []struct {
				Text string `xml:",chardata"`
				Name string `xml:"name,attr"`
			} `xml:"stringProp"`
			BoolProp []struct {
				Text string `xml:",chardata"`
				Name string `xml:"name,attr"`
			} `xml:"boolProp"`
			ElementProp struct {
				Text           string `xml:",chardata"`
				Name           string `xml:"name,attr"`
				ElementType    string `xml:"elementType,attr"`
				Guiclass       string `xml:"guiclass,attr"`
				Testclass      string `xml:"testclass,attr"`
				Enabled        string `xml:"enabled,attr"`
				CollectionProp struct {
					Text string `xml:",chardata"`
					Name string `xml:"name,attr"`
				} `xml:"collectionProp"`
			} `xml:"elementProp"`
		} `xml:"TestPlan"`
		HashTree struct {
			Text          string `xml:",chardata"`
			HeaderManager struct {
				Text           string `xml:",chardata"`
				Guiclass       string `xml:"guiclass,attr"`
				Testclass      string `xml:"testclass,attr"`
				Testname       string `xml:"testname,attr"`
				Enabled        string `xml:"enabled,attr"`
				CollectionProp struct {
					Text        string `xml:",chardata"`
					Name        string `xml:"name,attr"`
					ElementProp struct {
						Text        string `xml:",chardata"`
						Name        string `xml:"name,attr"`
						ElementType string `xml:"elementType,attr"`
						StringProp  []struct {
							Text string `xml:",chardata"`
							Name string `xml:"name,attr"`
						} `xml:"stringProp"`
					} `xml:"elementProp"`
				} `xml:"collectionProp"`
			} `xml:"HeaderManager"`
			HashTree []struct {
				Text                  string `xml:",chardata"`
				TransactionController struct {
					Text      string `xml:",chardata"`
					Guiclass  string `xml:"guiclass,attr"`
					Testclass string `xml:"testclass,attr"`
					Testname  string `xml:"testname,attr"`
					Enabled   string `xml:"enabled,attr"`
					BoolProp  []struct {
						Text string `xml:",chardata"`
						Name string `xml:"name,attr"`
					} `xml:"boolProp"`
				} `xml:"TransactionController"`
				HashTree struct {
					Text       string `xml:",chardata"`
					CSVDataSet struct {
						Text       string `xml:",chardata"`
						Guiclass   string `xml:"guiclass,attr"`
						Testclass  string `xml:"testclass,attr"`
						Testname   string `xml:"testname,attr"`
						Enabled    string `xml:"enabled,attr"`
						StringProp []struct {
							Text string `xml:",chardata"`
							Name string `xml:"name,attr"`
						} `xml:"stringProp"`
						BoolProp []struct {
							Text string `xml:",chardata"`
							Name string `xml:"name,attr"`
						} `xml:"boolProp"`
					} `xml:"CSVDataSet"`
					HashTree []struct {
						Text          string `xml:",chardata"`
						HeaderManager struct {
							Text           string `xml:",chardata"`
							Guiclass       string `xml:"guiclass,attr"`
							Testclass      string `xml:"testclass,attr"`
							Testname       string `xml:"testname,attr"`
							Enabled        string `xml:"enabled,attr"`
							CollectionProp struct {
								Text        string `xml:",chardata"`
								Name        string `xml:"name,attr"`
								ElementProp []struct {
									Text        string `xml:",chardata"`
									Name        string `xml:"name,attr"`
									ElementType string `xml:"elementType,attr"`
									StringProp  []struct {
										Text string `xml:",chardata"`
										Name string `xml:"name,attr"`
									} `xml:"stringProp"`
								} `xml:"elementProp"`
							} `xml:"collectionProp"`
						} `xml:"HeaderManager"`
						HashTree           []string `xml:"hashTree"`
						UniformRandomTimer struct {
							Text       string `xml:",chardata"`
							Guiclass   string `xml:"guiclass,attr"`
							Testclass  string `xml:"testclass,attr"`
							Testname   string `xml:"testname,attr"`
							Enabled    string `xml:"enabled,attr"`
							StringProp []struct {
								Text string `xml:",chardata"`
								Name string `xml:"name,attr"`
							} `xml:"stringProp"`
						} `xml:"UniformRandomTimer"`
						RegexExtractor []struct {
							Text       string `xml:",chardata"`
							Guiclass   string `xml:"guiclass,attr"`
							Testclass  string `xml:"testclass,attr"`
							Testname   string `xml:"testname,attr"`
							Enabled    string `xml:"enabled,attr"`
							StringProp []struct {
								Text string `xml:",chardata"`
								Name string `xml:"name,attr"`
							} `xml:"stringProp"`
						} `xml:"RegexExtractor"`
						XPathExtractor struct {
							Text       string `xml:",chardata"`
							Guiclass   string `xml:"guiclass,attr"`
							Testclass  string `xml:"testclass,attr"`
							Testname   string `xml:"testname,attr"`
							Enabled    string `xml:"enabled,attr"`
							StringProp []struct {
								Text string `xml:",chardata"`
								Name string `xml:"name,attr"`
							} `xml:"stringProp"`
							BoolProp []struct {
								Text string `xml:",chardata"`
								Name string `xml:"name,attr"`
							} `xml:"boolProp"`
						} `xml:"XPathExtractor"`
					} `xml:"hashTree"`
					URLRewritingModifier struct {
						Text       string `xml:",chardata"`
						Guiclass   string `xml:"guiclass,attr"`
						Testclass  string `xml:"testclass,attr"`
						Testname   string `xml:"testname,attr"`
						Enabled    string `xml:"enabled,attr"`
						StringProp struct {
							Text string `xml:",chardata"`
							Name string `xml:"name,attr"`
						} `xml:"stringProp"`
						BoolProp []struct {
							Text string `xml:",chardata"`
							Name string `xml:"name,attr"`
						} `xml:"boolProp"`
					} `xml:"URLRewritingModifier"`
					HTTPSamplerProxy []struct {
						Text        string `xml:",chardata"`
						Guiclass    string `xml:"guiclass,attr"`
						Testclass   string `xml:"testclass,attr"`
						Testname    string `xml:"testname,attr"`
						Enabled     string `xml:"enabled,attr"`
						ElementProp struct {
							Text           string `xml:",chardata"`
							Name           string `xml:"name,attr"`
							ElementType    string `xml:"elementType,attr"`
							Guiclass       string `xml:"guiclass,attr"`
							Testclass      string `xml:"testclass,attr"`
							Enabled        string `xml:"enabled,attr"`
							CollectionProp struct {
								Text        string `xml:",chardata"`
								Name        string `xml:"name,attr"`
								ElementProp []struct {
									Text        string `xml:",chardata"`
									Name        string `xml:"name,attr"`
									ElementType string `xml:"elementType,attr"`
									BoolProp    []struct {
										Text string `xml:",chardata"`
										Name string `xml:"name,attr"`
									} `xml:"boolProp"`
									StringProp []struct {
										Text string `xml:",chardata"`
										Name string `xml:"name,attr"`
									} `xml:"stringProp"`
								} `xml:"elementProp"`
							} `xml:"collectionProp"`
						} `xml:"elementProp"`
						StringProp []struct {
							Text string `xml:",chardata"`
							Name string `xml:"name,attr"`
						} `xml:"stringProp"`
						BoolProp []struct {
							Text string `xml:",chardata"`
							Name string `xml:"name,attr"`
						} `xml:"boolProp"`
					} `xml:"HTTPSamplerProxy"`
					ResultCollector []struct {
						Text      string `xml:",chardata"`
						Guiclass  string `xml:"guiclass,attr"`
						Testclass string `xml:"testclass,attr"`
						Testname  string `xml:"testname,attr"`
						Enabled   string `xml:"enabled,attr"`
						BoolProp  struct {
							Text string `xml:",chardata"`
							Name string `xml:"name,attr"`
						} `xml:"boolProp"`
						ObjProp struct {
							Text  string `xml:",chardata"`
							Name  string `xml:"name"`
							Value struct {
								Text                               string `xml:",chardata"`
								Class                              string `xml:"class,attr"`
								Time                               string `xml:"time"`
								Latency                            string `xml:"latency"`
								Timestamp                          string `xml:"timestamp"`
								Success                            string `xml:"success"`
								Label                              string `xml:"label"`
								Code                               string `xml:"code"`
								Message                            string `xml:"message"`
								ThreadName                         string `xml:"threadName"`
								DataType                           string `xml:"dataType"`
								Encoding                           string `xml:"encoding"`
								Assertions                         string `xml:"assertions"`
								Subresults                         string `xml:"subresults"`
								ResponseData                       string `xml:"responseData"`
								SamplerData                        string `xml:"samplerData"`
								XML                                string `xml:"xml"`
								FieldNames                         string `xml:"fieldNames"`
								ResponseHeaders                    string `xml:"responseHeaders"`
								RequestHeaders                     string `xml:"requestHeaders"`
								ResponseDataOnError                string `xml:"responseDataOnError"`
								SaveAssertionResultsFailureMessage string `xml:"saveAssertionResultsFailureMessage"`
								AssertionsResultsToSave            string `xml:"assertionsResultsToSave"`
								Bytes                              string `xml:"bytes"`
								SentBytes                          string `xml:"sentBytes"`
								URL                                string `xml:"url"`
								ThreadCounts                       string `xml:"threadCounts"`
								IdleTime                           string `xml:"idleTime"`
								ConnectTime                        string `xml:"connectTime"`
							} `xml:"value"`
						} `xml:"objProp"`
						StringProp struct {
							Text string `xml:",chardata"`
							Name string `xml:"name,attr"`
						} `xml:"stringProp"`
					} `xml:"ResultCollector"`
				} `xml:"hashTree"`
			} `xml:"hashTree"`
			Arguments struct {
				Text           string `xml:",chardata"`
				Guiclass       string `xml:"guiclass,attr"`
				Testclass      string `xml:"testclass,attr"`
				Testname       string `xml:"testname,attr"`
				Enabled        string `xml:"enabled,attr"`
				CollectionProp struct {
					Text        string `xml:",chardata"`
					Name        string `xml:"name,attr"`
					ElementProp []struct {
						Text        string `xml:",chardata"`
						Name        string `xml:"name,attr"`
						ElementType string `xml:"elementType,attr"`
						StringProp  []struct {
							Text string `xml:",chardata"`
							Name string `xml:"name,attr"`
						} `xml:"stringProp"`
					} `xml:"elementProp"`
				} `xml:"collectionProp"`
			} `xml:"Arguments"`
			ConfigTestElement struct {
				Text        string `xml:",chardata"`
				Guiclass    string `xml:"guiclass,attr"`
				Testclass   string `xml:"testclass,attr"`
				Testname    string `xml:"testname,attr"`
				Enabled     string `xml:"enabled,attr"`
				ElementProp struct {
					Text           string `xml:",chardata"`
					Name           string `xml:"name,attr"`
					ElementType    string `xml:"elementType,attr"`
					Guiclass       string `xml:"guiclass,attr"`
					Testclass      string `xml:"testclass,attr"`
					Enabled        string `xml:"enabled,attr"`
					CollectionProp struct {
						Text string `xml:",chardata"`
						Name string `xml:"name,attr"`
					} `xml:"collectionProp"`
				} `xml:"elementProp"`
				StringProp []struct {
					Text string `xml:",chardata"`
					Name string `xml:"name,attr"`
				} `xml:"stringProp"`
				BoolProp []struct {
					Text string `xml:",chardata"`
					Name string `xml:"name,attr"`
				} `xml:"boolProp"`
			} `xml:"ConfigTestElement"`
			DNSCacheManager struct {
				Text           string `xml:",chardata"`
				Guiclass       string `xml:"guiclass,attr"`
				Testclass      string `xml:"testclass,attr"`
				Testname       string `xml:"testname,attr"`
				Enabled        string `xml:"enabled,attr"`
				CollectionProp struct {
					Text string `xml:",chardata"`
					Name string `xml:"name,attr"`
				} `xml:"collectionProp"`
				BoolProp []struct {
					Text string `xml:",chardata"`
					Name string `xml:"name,attr"`
				} `xml:"boolProp"`
			} `xml:"DNSCacheManager"`
			AuthManager struct {
				Text           string `xml:",chardata"`
				Guiclass       string `xml:"guiclass,attr"`
				Testclass      string `xml:"testclass,attr"`
				Testname       string `xml:"testname,attr"`
				Enabled        string `xml:"enabled,attr"`
				CollectionProp struct {
					Text string `xml:",chardata"`
					Name string `xml:"name,attr"`
				} `xml:"collectionProp"`
				BoolProp struct {
					Text string `xml:",chardata"`
					Name string `xml:"name,attr"`
				} `xml:"boolProp"`
			} `xml:"AuthManager"`
			CookieManager struct {
				Text           string `xml:",chardata"`
				Guiclass       string `xml:"guiclass,attr"`
				Testclass      string `xml:"testclass,attr"`
				Testname       string `xml:"testname,attr"`
				Enabled        string `xml:"enabled,attr"`
				CollectionProp struct {
					Text string `xml:",chardata"`
					Name string `xml:"name,attr"`
				} `xml:"collectionProp"`
				BoolProp []struct {
					Text string `xml:",chardata"`
					Name string `xml:"name,attr"`
				} `xml:"boolProp"`
			} `xml:"CookieManager"`
			CacheManager struct {
				Text      string `xml:",chardata"`
				Guiclass  string `xml:"guiclass,attr"`
				Testclass string `xml:"testclass,attr"`
				Testname  string `xml:"testname,attr"`
				Enabled   string `xml:"enabled,attr"`
				BoolProp  []struct {
					Text string `xml:",chardata"`
					Name string `xml:"name,attr"`
				} `xml:"boolProp"`
			} `xml:"CacheManager"`
			ThreadGroup struct {
				Text       string `xml:",chardata"`
				Guiclass   string `xml:"guiclass,attr"`
				Testclass  string `xml:"testclass,attr"`
				Testname   string `xml:"testname,attr"`
				Enabled    string `xml:"enabled,attr"`
				StringProp []struct {
					Text string `xml:",chardata"`
					Name string `xml:"name,attr"`
				} `xml:"stringProp"`
				ElementProp struct {
					Text        string `xml:",chardata"`
					Name        string `xml:"name,attr"`
					ElementType string `xml:"elementType,attr"`
					Guiclass    string `xml:"guiclass,attr"`
					Testclass   string `xml:"testclass,attr"`
					Enabled     string `xml:"enabled,attr"`
					BoolProp    struct {
						Text string `xml:",chardata"`
						Name string `xml:"name,attr"`
					} `xml:"boolProp"`
					StringProp struct {
						Text string `xml:",chardata"`
						Name string `xml:"name,attr"`
					} `xml:"stringProp"`
				} `xml:"elementProp"`
				BoolProp []struct {
					Text string `xml:",chardata"`
					Name string `xml:"name,attr"`
				} `xml:"boolProp"`
			} `xml:"ThreadGroup"`
		} `xml:"hashTree"`
	} `xml:"hashTree"`
}
