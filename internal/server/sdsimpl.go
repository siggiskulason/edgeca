/*******************************************************************************
 * Copyright 2021 EdgeSec OÜ
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 *
 *******************************************************************************/

package server

import (
	"context"
	"crypto/x509/pkix"
	"os"

	"github.com/edgesec-org/edgeca/internal/certs"
	"github.com/edgesec-org/edgeca/internal/server/state"
	core "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	tls "github.com/envoyproxy/go-control-plane/envoy/extensions/transport_sockets/tls/v3"
	secretservice "github.com/envoyproxy/go-control-plane/envoy/service/secret/v3"
	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	sdsserver "github.com/envoyproxy/go-control-plane/pkg/server/v3"
	test "github.com/envoyproxy/go-control-plane/pkg/test/v3"
	"google.golang.org/grpc"
)

func generateSDSCertificate() (pemCert, pemKey string, err error) {

	pemCertificate, pemPrivateKey, err := certs.GeneratePemCertificate(pkix.Name{CommonName: "localhost"}, state.GetSubCACert(), state.GetSubCAKey())
	return string(pemCertificate), string(pemPrivateKey), err
}

func makeSecret() *tls.Secret {

	pemCert, pemKey, _ := generateSDSCertificate()

	certificate := &tls.Secret_TlsCertificate{}
	certificate.TlsCertificate = &tls.TlsCertificate{}
	certificate.TlsCertificate.CertificateChain = &core.DataSource{}
	certificate.TlsCertificate.CertificateChain.Specifier = &core.DataSource_InlineString{
		InlineString: pemCert,
	}

	certificate.TlsCertificate.PrivateKey = &core.DataSource{}
	certificate.TlsCertificate.PrivateKey.Specifier = &core.DataSource_InlineString{
		InlineString: pemKey,
	}

	result := &tls.Secret{}
	result.Name = "server_cert"
	result.Type = certificate

	return result
}

func GenerateSnapshot() cache.Snapshot {
	return cache.NewSnapshot(
		"1",
		[]types.Resource{},
		[]types.Resource{},
		[]types.Resource{},
		[]types.Resource{},
		[]types.Resource{},
		[]types.Resource{makeSecret()},
	)
}

func StartSDSServer(port int, grpcServer *grpc.Server) {

	cache := cache.NewSnapshotCache(false, cache.IDHash{}, nil)

	snapshot := GenerateSnapshot()
	if err := snapshot.Consistent(); err != nil {
		os.Exit(1)
	}

	if err := cache.SetSnapshot("test", snapshot); err != nil {
		os.Exit(1)
	}

	ctx := context.Background()
	cb := &test.Callbacks{Debug: true}
	srv := sdsserver.NewServer(ctx, cache, cb)

	secretservice.RegisterSecretDiscoveryServiceServer(grpcServer, srv)

}
