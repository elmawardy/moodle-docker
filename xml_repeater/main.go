package main

import (
	"fmt"
)

func main(){
	httpSampler := `          <hashTree>
            <UniformRandomTimer guiclass="UniformRandomTimerGui" testclass="UniformRandomTimer" testname="Uniform Random Timer" enabled="true">
              <stringProp name="ConstantTimer.delay">2000</stringProp>
              <stringProp name="RandomTimer.range">30000</stringProp>
            </UniformRandomTimer>
            <hashTree/>
          </hashTree>
          <HTTPSamplerProxy guiclass="HttpTestSampleGui" testclass="HTTPSamplerProxy" testname="/mod/quiz/processattempt.php?cmid=${quiz_id}" enabled="true">
            <elementProp name="HTTPsampler.Arguments" elementType="Arguments" guiclass="HTTPArgumentsPanel" testclass="Arguments" enabled="true">
              <collectionProp name="Arguments.arguments">
                <elementProp name="cmid" elementType="HTTPArgument">
                  <boolProp name="HTTPArgument.always_encode">false</boolProp>
                  <stringProp name="Argument.name">cmid</stringProp>
                  <stringProp name="Argument.value">${quiz_id}</stringProp>
                  <stringProp name="Argument.metadata">=</stringProp>
                  <boolProp name="HTTPArgument.use_equals">true</boolProp>
                </elementProp>
                <elementProp name="next" elementType="HTTPArgument">
                  <boolProp name="HTTPArgument.always_encode">true</boolProp>
                  <stringProp name="Argument.name">next</stringProp>
                  <stringProp name="Argument.value">Next page</stringProp>
                  <stringProp name="Argument.metadata">=</stringProp>
                  <boolProp name="HTTPArgument.use_equals">true</boolProp>
                </elementProp>
                <elementProp name="q${useage_id}:%d_answer" elementType="HTTPArgument">
                  <boolProp name="HTTPArgument.always_encode">true</boolProp>
                  <stringProp name="Argument.name">q${useage_id}:%d_answer</stringProp>
                  <stringProp name="Argument.value">${__Random(0,1)}</stringProp>
                  <stringProp name="Argument.metadata">=</stringProp>
                  <boolProp name="HTTPArgument.use_equals">true</boolProp>
                </elementProp>
                <elementProp name="nextpage" elementType="HTTPArgument">
                  <boolProp name="HTTPArgument.always_encode">true</boolProp>
                  <stringProp name="Argument.name">nextpage</stringProp>
                  <stringProp name="Argument.value">%d</stringProp>
                  <stringProp name="Argument.metadata">=</stringProp>
                  <boolProp name="HTTPArgument.use_equals">true</boolProp>
                </elementProp>
                <elementProp name="q${useage_id}:%d_:flagged" elementType="HTTPArgument">
                  <boolProp name="HTTPArgument.always_encode">true</boolProp>
                  <stringProp name="Argument.name">q${useage_id}:%d_:flagged</stringProp>
                  <stringProp name="Argument.value">0</stringProp>
                  <stringProp name="Argument.metadata">=</stringProp>
                  <boolProp name="HTTPArgument.use_equals">true</boolProp>
                </elementProp>
                <elementProp name="slots" elementType="HTTPArgument">
                  <boolProp name="HTTPArgument.always_encode">true</boolProp>
                  <stringProp name="Argument.name">slots</stringProp>
                  <stringProp name="Argument.value">%d</stringProp>
                  <stringProp name="Argument.metadata">=</stringProp>
                  <boolProp name="HTTPArgument.use_equals">true</boolProp>
                </elementProp>
                <elementProp name="timeup" elementType="HTTPArgument">
                  <boolProp name="HTTPArgument.always_encode">true</boolProp>
                  <stringProp name="Argument.name">timeup</stringProp>
                  <stringProp name="Argument.value">0</stringProp>
                  <stringProp name="Argument.metadata">=</stringProp>
                  <boolProp name="HTTPArgument.use_equals">true</boolProp>
                </elementProp>
                <elementProp name="q${useage_id}:%d_:sequencecheck" elementType="HTTPArgument">
                  <boolProp name="HTTPArgument.always_encode">true</boolProp>
                  <stringProp name="Argument.name">q${useage_id}:%d_:sequencecheck</stringProp>
                  <stringProp name="Argument.value">1</stringProp>
                  <stringProp name="Argument.metadata">=</stringProp>
                  <boolProp name="HTTPArgument.use_equals">true</boolProp>
                </elementProp>
                <elementProp name="thispage" elementType="HTTPArgument">
                  <boolProp name="HTTPArgument.always_encode">true</boolProp>
                  <stringProp name="Argument.name">thispage</stringProp>
                  <stringProp name="Argument.value">%d</stringProp>
                  <stringProp name="Argument.metadata">=</stringProp>
                  <boolProp name="HTTPArgument.use_equals">true</boolProp>
                </elementProp>
                <elementProp name="attempt" elementType="HTTPArgument">
                  <boolProp name="HTTPArgument.always_encode">true</boolProp>
                  <stringProp name="Argument.name">attempt</stringProp>
                  <stringProp name="Argument.value">${attempt_id}</stringProp>
                  <stringProp name="Argument.metadata">=</stringProp>
                  <boolProp name="HTTPArgument.use_equals">true</boolProp>
                </elementProp>
                <elementProp name="sesskey" elementType="HTTPArgument">
                  <boolProp name="HTTPArgument.always_encode">true</boolProp>
                  <stringProp name="Argument.name">sesskey</stringProp>
                  <stringProp name="Argument.value">${sesskey}</stringProp>
                  <stringProp name="Argument.metadata">=</stringProp>
                  <boolProp name="HTTPArgument.use_equals">true</boolProp>
                </elementProp>
              </collectionProp>
            </elementProp>
            <stringProp name="HTTPSampler.domain">${BASE_URL_1}</stringProp>
            <stringProp name="HTTPSampler.port"></stringProp>
            <stringProp name="HTTPSampler.protocol">https</stringProp>
            <stringProp name="HTTPSampler.contentEncoding"></stringProp>
            <stringProp name="HTTPSampler.path">mod/quiz/processattempt.php</stringProp>
            <stringProp name="HTTPSampler.method">POST</stringProp>
            <boolProp name="HTTPSampler.follow_redirects">true</boolProp>
            <boolProp name="HTTPSampler.auto_redirects">false</boolProp>
            <boolProp name="HTTPSampler.use_keepalive">true</boolProp>
            <boolProp name="HTTPSampler.DO_MULTIPART_POST">false</boolProp>
            <stringProp name="HTTPSampler.embedded_url_re"></stringProp>
            <stringProp name="HTTPSampler.connect_timeout"></stringProp>
            <stringProp name="HTTPSampler.response_timeout"></stringProp>
          </HTTPSamplerProxy>
`;



        for i:=1;i<=100;i++{
		fmt.Printf(httpSampler,i,i,i,i,i,i,i,i,i-1)
	}



}
