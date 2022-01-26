# V1Probe

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exec** | [***V1ExecAction**](V1ExecAction.md) |  | [optional] [default to null]
**FailureThreshold** | **int32** | Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1. | [optional] [default to null]
**HttpGet** | [***V1HttpGetAction**](V1HTTPGetAction.md) |  | [optional] [default to null]
**InitialDelaySeconds** | **int32** | Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes | [optional] [default to null]
**PeriodSeconds** | **int32** | How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1. | [optional] [default to null]
**SuccessThreshold** | **int32** | Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1. | [optional] [default to null]
**TcpSocket** | [***V1TcpSocketAction**](V1TCPSocketAction.md) |  | [optional] [default to null]
**TimeoutSeconds** | **int32** | Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


