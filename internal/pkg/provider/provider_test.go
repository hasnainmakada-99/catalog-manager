/**
 * Copyright 2021 Napptive
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package provider

import (
	"github.com/napptive/catalog-manager/internal/pkg/entities"
	"github.com/napptive/catalog-manager/internal/pkg/utils"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"time"
)

func RunTests(provider MetadataProvider) {

	ginkgo.Context("Adding application metadata", func() {
		ginkgo.It("Should be able to add an application metadata", func() {
			app := utils.CreateApplicationMetadata()

			returned, err := provider.Add(app)
			gomega.Expect(err).Should(gomega.Succeed())
			gomega.Expect(returned.CatalogID).ShouldNot(gomega.BeEmpty())

		})

		ginkgo.It("Should be able to add an application metadata twice (update)", func() {
			app := utils.CreateApplicationMetadata()

			_, err := provider.Add(app)
			gomega.Expect(err).Should(gomega.Succeed())

			_, err = provider.Add(app)
			gomega.Expect(err).Should(gomega.Succeed())
		})

	})

	ginkgo.Context("Getting application metadata", func() {

		ginkgo.It("Should be able to get an application metadata", func() {
			app := utils.CreateApplicationMetadata()

			returned, err := provider.Add(app)
			gomega.Expect(err).Should(gomega.Succeed())

			// wait to be stored
			time.Sleep(time.Second)

			retrieved, err := provider.Get(entities.ApplicationID{
				Repository:      returned.Repository,
				ApplicationName: returned.ApplicationName,
				Tag:             returned.Tag,
			})
			gomega.Expect(err).Should(gomega.Succeed())
			gomega.Expect(retrieved).ShouldNot(gomega.BeNil())
			gomega.Expect(*retrieved).Should(gomega.Equal(*app))
		})

		ginkgo.It("Should not be able to get an application metadata if it does not exist", func() {
			_, err := provider.Get(entities.ApplicationID{
				Repository:      "repoTest",
				ApplicationName: "applName",
				Tag:             "",
			})
			gomega.Expect(err).ShouldNot(gomega.Succeed())
		})

	})

	ginkgo.Context("Removing application metadata", func() {
		ginkgo.It("Should be able to delete an application metadata", func() {
			app := utils.CreateApplicationMetadata()

			returned, err := provider.Add(app)
			gomega.Expect(err).Should(gomega.Succeed())

			// wait to be stored
			time.Sleep(time.Second)

			err = provider.Remove(&entities.ApplicationID{
				Repository:      returned.Repository,
				ApplicationName: returned.ApplicationName,
				Tag:             returned.Tag,
			})
			gomega.Expect(err).Should(gomega.Succeed())

		})
		ginkgo.It("Should not be able to delete an application metadata if it does not exist", func() {
			app := utils.CreateApplicationMetadata()

			err := provider.Remove(&entities.ApplicationID{
				Repository:      app.Repository,
				ApplicationName: app.ApplicationName,
				Tag:             app.Tag,
			})
			gomega.Expect(err).ShouldNot(gomega.Succeed())

		})
	})

}