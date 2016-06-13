package vast_test

import (
	"github.com/Vungle/vungo/vast"
	"github.com/Vungle/vungo/vast/vasttest"
	"reflect"
	"testing"
)

var CreativeWrapperModelType = reflect.TypeOf(vast.CreativeWrapper{})

func TestCreativeWrapperMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "Creative", "creativewrapper.xml", CreativeWrapperModelType)
}

var creativeWrapperTests = []vasttest.VastTest{
	vasttest.VastTest{&vast.CreativeWrapper{}, vast.ErrCreativeWrapperType, "creativewrapper.xml"},
	vasttest.VastTest{&vast.CreativeWrapper{}, nil, "creativewrapper_with_linear.xml"},
	vasttest.VastTest{&vast.CreativeWrapper{}, nil, "creativewrapper_with_nonlinearads.xml"},
	vasttest.VastTest{&vast.CreativeWrapper{}, nil, "creativewrapper_with_companionads.xml"},
	vasttest.VastTest{&vast.CreativeWrapper{}, vast.ErrCreativeWrapperType, "creativewrapper_without_linear.xml"},
	vasttest.VastTest{&vast.CreativeWrapper{}, vast.ErrCreativeWrapperType, "creativewrapper_without_nonlinearads.xml"},
	vasttest.VastTest{&vast.CreativeWrapper{}, vast.ErrCreativeWrapperType, "creativewrapper_without_companionads.xml"},
	vasttest.VastTest{&vast.CreativeWrapper{}, vast.ErrIconResourcesFormat, "creativewrapper_error_linear.xml"},
	vasttest.VastTest{&vast.CreativeWrapper{}, vast.ErrTrackingMissUri, "creativewrapper_error_nonlinearads.xml"},
	vasttest.VastTest{&vast.CreativeWrapper{}, vast.ErrCompanionAdsWrapperMissCompanions, "creativewrapper_error_companionads.xml"},
	vasttest.VastTest{&vast.CreativeWrapper{}, vast.ErrCreativeWrapperType, "creativewrapper_without_ads.xml"},
}

func TestCreativeWrapperValidateErrors(t *testing.T) {
	for _, test := range creativeWrapperTests {
		vasttest.VerifyVastElementFromFile(t, "testdata/"+test.File, test.VastElement, test.Err)
	}
}
