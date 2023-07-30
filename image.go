package image

func CWRotate90(targetImagePath, outputImagePath string) error {
	return New(targetImagePath).CWRotate90(outputImagePath)
}

func CWRotate180(targetImagePath, outputImagePath string) error {
	return New(targetImagePath).CWRotate180(outputImagePath)
}

func CWRotate270(targetImagePath, outputImagePath string) error {
	return New(targetImagePath).CWRotate270(outputImagePath)
}
