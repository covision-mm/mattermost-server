// Copyright (c) 2016-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

// Code generated by enterprise/einterfaces_gen.go; DO NOT EDIT.

package einterfaces

import (
	"github.com/gorilla/mux"
)

type MetricsInterface interface {
	AttachProfiler(router *mux.Router)
	StartServer()
	StopServer()
	IncrementPostCreate()
	IncrementWebhookPost()
	IncrementPostSentEmail()
	IncrementPostSentPush()
	IncrementPostBroadcast()
	IncrementWebSocketBroadcast(eventType string)
	IncrementPostFileAttachment(count int)
	IncrementHttpRequest()
	IncrementHttpError()
	ObserveHttpRequestDuration(elapsed float64)
	IncrementClusterRequest()
	ObserveClusterRequestDuration(elapsed float64)
	IncrementClusterEventType(eventType string)
	IncrementLogin()
	IncrementLoginFail()
	IncrementEtagMissCounter(route string)
	IncrementEtagHitCounter(route string)
	IncrementMemCacheMissCounter(cacheName string)
	IncrementMemCacheHitCounter(cacheName string)
	IncrementMemCacheInvalidationCounter(cacheName string)
	IncrementMemCacheMissCounterSession()
	IncrementMemCacheHitCounterSession()
	IncrementMemCacheInvalidationCounterSession()
	AddMemCacheMissCounter(cacheName string, amount float64)
	AddMemCacheHitCounter(cacheName string, amount float64)
	IncrementWebsocketEvent(eventType string)
	IncrementPostsSearchCounter()
	ObservePostsSearchDuration(elapsed float64)
}
