package fake

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	gojenkins "github.com/jenkins-x/golang-jenkins"
)

// NotFoundMessage error message if something is not found
const NotFoundMessage = "not found"

// FakeClient a fake Jenkins client for easier testing
type Client struct {
	Jobs          []gojenkins.Job
	BaseURLValue  string
	Operations    []string
	XMLJobs       []XMLJob
	FolderXMLJobs []FolderXMLJob
	BuildRequests []BuildRequest

	httpClient *http.Client
}

// XMLJob represents a fake created XML Job
type XMLJob struct {
	JobItemXml string
	JobName    string
}

// FolderXMLJob represents a fake created folder XML Job
type FolderXMLJob struct {
	JobItemXml string
	Folder     string
	JobName    string
}

// BuildRequest for recording build requests
type BuildRequest struct {
	Job    gojenkins.Job
	Values url.Values
}

// FullJobPath returns the full job path
func (j *FolderXMLJob) FullJobPath() string {
	return "/job/" + j.Folder + "/job/" + j.JobName
}

var _ gojenkins.JenkinsClient = (*Client)(nil)

func (f *Client) GetJobs() ([]gojenkins.Job, error) {
	return f.Jobs, nil
}

func (f *Client) GetJob(name string) (gojenkins.Job, error) {
	for _, j := range f.Jobs {
		if j.Name == name {
			return j, nil
		}
	}
	return gojenkins.Job{}, notFoundError()
}

func notFoundError() error {
	return fmt.Errorf(NotFoundMessage)
}

func (f *Client) GetJobURLPath(path string) string {
	panic("implement me")
}

func (f *Client) IsErrNotFound(err error) bool {
	return err.Error() == NotFoundMessage
}

func (f *Client) BaseURL() string {
	return f.BaseURLValue
}

func (f *Client) SetHTTPClient(httpClient *http.Client) {
	f.httpClient = httpClient
}

func (f *Client) Post(string, url.Values, interface{}) (err error) {
	panic("implement me")
}

func (f *Client) GetJobConfig(string) (gojenkins.JobItem, error) {
	panic("implement me")
}

func (f *Client) GetBuild(gojenkins.Job, int) (gojenkins.Build, error) {
	panic("implement me")
}

func (f *Client) GetLastBuild(gojenkins.Job) (gojenkins.Build, error) {
	panic("implement me")
}

func (f *Client) StopBuild(gojenkins.Job, int) error {
	panic("implement me")
}

func (f *Client) GetMultiBranchJob(string, string, string) (gojenkins.Job, error) {
	panic("implement me")
}

func (f *Client) GetJobByPath(paths ...string) (gojenkins.Job, error) {
	var job gojenkins.Job

	fullPath := gojenkins.FullJobPath(paths...)
	for _, xj := range f.XMLJobs {
		if xj.JobName == fullPath {
			job.Name = xj.JobName
			job.FullName = fullPath
			return job, nil
		}
	}
	for _, xj := range f.FolderXMLJobs {
		fullFolderPath := xj.FullJobPath()
		// the fullFolderPath won't have the /job/master prefix
		if strings.HasPrefix(fullFolderPath, fullPath) {
			job.Name = xj.JobName
			job.FullName = fullPath
			return job, nil
		}
	}

	var err error
	for i, name := range paths {
		if i == 0 {
			job, err = f.GetJob(name)
			if err != nil {
				return job, err
			}

		} else {
			for _, j := range job.Jobs {
				if j.Name == name {
					job = j
					continue
				}
			}
			return gojenkins.Job{}, notFoundError()
		}
	}
	return gojenkins.Job{}, notFoundError()
}

func (f *Client) GetOrganizationScanResult(int, gojenkins.Job) (string, error) {
	panic("implement me")
}

func (f *Client) CreateJob(gojenkins.JobItem, string) error {
	panic("implement me")
}

func (f *Client) Reload() error {
	return f.doOperation("Reload")
}

func (f *Client) Restart() error {
	return f.doOperation("Restart")
}

func (f *Client) SafeRestart() error {
	return f.doOperation("SafeRestart")
}

func (f *Client) QuietDown() error {
	return f.doOperation("QuietDown")
}

func (f *Client) CreateJobWithXML(jobItemXml string, jobName string) error {
	f.XMLJobs = append(f.XMLJobs, XMLJob{jobItemXml, jobName})
	return nil
}

func (f *Client) CreateFolderJobWithXML(jobItemXml string, folder string, jobName string) error {
	f.FolderXMLJobs = append(f.FolderXMLJobs, FolderXMLJob{jobItemXml, folder, jobName})
	return nil
}

func (f *Client) GetCredential(string) (*gojenkins.Credentials, error) {
	panic("implement me")
}

func (f *Client) CreateCredential(string, string, string) error {
	panic("implement me")
}

func (f *Client) DeleteJob(gojenkins.Job) error {
	panic("implement me")
}

func (f *Client) UpdateJob(gojenkins.JobItem, string) error {
	panic("implement me")
}

func (f *Client) RemoveJob(string) error {
	panic("implement me")
}

func (f *Client) AddJobToView(string, gojenkins.Job) error {
	panic("implement me")
}

func (f *Client) CreateView(gojenkins.ListView) error {
	panic("implement me")
}

func (f *Client) Build(job gojenkins.Job, values url.Values) error {
	f.BuildRequests = append(f.BuildRequests, BuildRequest{job, values})
	return nil
}

func (f *Client) GetBuildConsoleOutput(gojenkins.Build) ([]byte, error) {
	panic("implement me")
}

func (f *Client) GetQueue() (gojenkins.Queue, error) {
	panic("implement me")
}

func (f *Client) GetArtifact(gojenkins.Build, gojenkins.Artifact) ([]byte, error) {
	panic("implement me")
}

func (f *Client) SetBuildDescription(gojenkins.Build, string) error {
	panic("implement me")
}

func (f *Client) GetComputerObject() (gojenkins.ComputerObject, error) {
	panic("implement me")
}

func (f *Client) GetComputers() ([]gojenkins.Computer, error) {
	panic("implement me")
}

func (f *Client) GetComputer(string) (gojenkins.Computer, error) {
	panic("implement me")
}

func (f *Client) GetBuildURL(gojenkins.Job, int) string {
	panic("implement me")
}

func (f *Client) GetLogFromURL(string, int64, *gojenkins.LogData) error {
	panic("implement me")
}

func (f *Client) TailLog(string, io.Writer, time.Duration, time.Duration) error {
	panic("implement me")
}

func (f *Client) TailLogFunc(string, io.Writer) gojenkins.ConditionFunc {
	panic("implement me")
}

func (f *Client) NewLogPoller(string, io.Writer) *gojenkins.LogPoller {
	panic("implement me")
}

func (f *Client) doOperation(name string) error {
	f.Operations = append(f.Operations, name)
	return nil
}
