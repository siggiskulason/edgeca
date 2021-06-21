package edgeca
		

	csr_policy {

        re_match(`^Darval Solutions Ltd$`, input.csr.Subject.Organization[0])

        # input.csr.Subject.Organization == organization
    }
