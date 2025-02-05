// Tencent is pleased to support the open source community by making
// 蓝鲸智云 - 监控平台 (BlueKing - Monitor) available.
// Copyright (C) 2022 THL A29 Limited, a Tencent company. All rights reserved.
// Licensed under the MIT License (the "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://opensource.org/licenses/MIT
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
// an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.

// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

//go:build SkipList
// +build SkipList

package storage_test

import "testing"

// BenchmarkStoreSet_SkipList :
func BenchmarkStoreSet_SkipList(b *testing.B) {
	withClosingStore(benchmarkStoreSet, b, newSkipList())
}

// BenchmarkStoreUpdate_SkipList :
func BenchmarkStoreUpdate_SkipList(b *testing.B) {
	withClosingStore(benchmarkStoreUpdate, b, newSkipList())
}

// BenchmarkStoreGet_SkipList :
func BenchmarkStoreGet_SkipList(b *testing.B) {
	withClosingStore(benchmarkStoreGet, b, newSkipList())
}

// BenchmarkStoreGetHotPot_SkipList :
func BenchmarkStoreGetHotPot_SkipList(b *testing.B) {
	withClosingStore(benchmarkStoreGetHotPot, b, newSkipList())
}

// benchmarkStoreExistsMissing_SkipList :
func BenchmarkStoreExistsMissing_SkipList(b *testing.B) {
	withClosingStore(benchmarkStoreExistsMissing, b, newSkipList())
}

// BenchmarkStoreExists_SkipList :
func BenchmarkStoreExists_SkipList(b *testing.B) {
	withClosingStore(benchmarkStoreExists, b, newSkipList())
}

// BenchmarkStoreDelete_SkipList :
func BenchmarkStoreDelete_SkipList(b *testing.B) {
	withClosingStore(benchmarkStoreDelete, b, newSkipList())
}

// BenchmarkStoreScan_SkipList :
func BenchmarkStoreScan_SkipList(b *testing.B) {
	withClosingStore(benchmarkStoreScan, b, newSkipList())
}

// BenchmarkStoreCommit_SkipList :
func BenchmarkStoreCommit_SkipList(b *testing.B) {
	withClosingStore(benchmarkStoreCommit, b, newSkipList())
}
