package helper

import (
	"strings"
)

func ErrorRequestMessage(err error) string {
	var message = "Error : "
	errors := []DuplicateError{
		// DISCOUNT PROPOSAL CONFIRMATION
		{
			contains: "DiscountProposalConfirmationCreateRequest.DiscountProposalID",
			message:  "Field validation for discount_proposal_id required, max 15 character",
		},
		{
			contains: "DiscountProposalConfirmationCreateRequest.DiscountProposalConfirmationStatusStep",
			message:  "Field validation for discount_proposal_confirmation_status_step required",
		},
		{
			contains: "DiscountProposalConfirmationCreateRequest.DiscountProposalConfirmationStatusType",
			message:  "Field validation for discount_proposal_confirmation_status_type required",
		},
		{
			contains: "DiscountProposalConfirmationCreateRequest.Note",
			message:  "Field validation for note required, max 500 character",
		},

		// DISCOUNT PROPOSAL CONFIRMATION STATUS
		{
			contains: "DiscountProposalConfirmationStatusCreateRequest.Name",
			message:  "Field validation for name required, max 30 character",
		},
		{
			contains: "DiscountProposalConfirmationStatusCreateRequest.Dateline",
			message:  "Field validation for dateline required, gte = 1 lte = 31",
		},
		{
			contains: "DiscountProposalConfirmationStatusCreateRequest.Unit",
			message:  "Field validation for unit required, value is not listed",
		},
		{
			contains: "DiscountProposalConfirmationStatusCreateRequest.Type",
			message:  "Field validation for type required, value is not listed",
		},
		{
			contains: "DiscountProposalConfirmationStatusCreateRequest.Step",
			message:  "Field validation for note required, lte = 10 ",
		},
		{
			contains: "DiscountProposalConfirmationStatusUpdateRequest.Name",
			message:  "Field validation for name required, max 30 character",
		},
		{
			contains: "DiscountProposalConfirmationStatusUpdateRequest.Dateline",
			message:  "Field validation for dateline required, gte = 1 lte = 31",
		},
		{
			contains: "DiscountProposalConfirmationStatusUpdateRequest.Unit",
			message:  "Field validation for unit required, value is not listed",
		},
		{
			contains: "DiscountProposalConfirmationStatusUpdateRequest.Type",
			message:  "Field validation for type required, value is not listed",
		},
		{
			contains: "DiscountProposalConfirmationStatusUpdateRequest.Step",
			message:  "Field validation for note required, lte = 10 ",
		},

		// DISCOUNT PROPOSAL
		{
			contains: "DiscountProposalCreateRequest.DivisionID",
			message:  "Field validation for division_id required",
		},
		{
			contains: "DiscountProposalCreateRequest.MarketingStructureID",
			message:  "Field validation for marketing_structure_id required",
		},
		{
			contains: "DiscountProposalCreateRequest.CityID",
			message:  "Field validation for city_id required",
		},
		{
			contains: "DiscountProposalCreateRequest.Period",
			message:  "Field validation for period required, format is wrong (YYYYMM)",
		},
		{
			contains: "DiscountProposalCreateRequest.PeriodStart",
			message:  "Field validation for period_start required, format is wrong (YYYYMMDD)",
		},
		{
			contains: "DiscountProposalCreateRequest.PeriodEnd",
			message:  "Field validation for period_end required, format is wrong (YYYYMMDD)",
		},
		{
			contains: "DiscountProposalCreateRequest.Type",
			message:  "Field validation for type required",
		},
		{
			contains: "DiscountProposalUpdateRequest.DivisionID",
			message:  "Field validation for division_id required",
		},
		{
			contains: "DiscountProposalUpdateRequest.MarketingStructureID",
			message:  "Field validation for marketing_structure_id required",
		},
		{
			contains: "DiscountProposalUpdateRequest.CityID",
			message:  "Field validation for city_id required",
		},
		{
			contains: "DiscountProposalUpdateRequest.Period",
			message:  "Field validation for period required, format is wrong (YYYYMM)",
		},
		{
			contains: "DiscountProposalUpdateRequest.PeriodStart",
			message:  "Field validation for period_start required, format is wrong (YYYYMMDD)",
		},
		{
			contains: "DiscountProposalUpdateRequest.PeriodEnd",
			message:  "Field validation for period_end required, format is wrong (YYYYMMDD)",
		},
		{
			contains: "DiscountProposalUpdateRequest.Type",
			message:  "Field validation for type required",
		},

		// CREATE RANGKUMAN
		{
			contains: "DiscountProposalCreateRangkumanRequest.DiscountProposalID",
			message:  "Field validation for discount_proposal_id required",
		},

		// CREATE MEMO
		{
			contains: "DiscountProposalCreateMemoRequest.DiscountProposalID",
			message:  "Field validation for discount_proposal_id required",
		},
		{
			contains: "DiscountProposalCreateMemoRequest.RangkumanNo",
			message:  "Field validation for rangkuman_no required",
		},

		// DISCOUNT PROPOSAL ESTIMATION
		{
			contains: "DiscountProposalEstimationCreateRequest.DiscountProposalID",
			message:  "Field validation for discount_proposal_id required",
		},
		{
			contains: "DiscountProposalEstimationCreateRequest.OutletID",
			message:  "Field validation for outlet_id required",
		},
		{
			contains: "DiscountProposalEstimationCreateRequest.OutletInvoiceID",
			message:  "Field validation for outlet_invoice_id required",
		},
		{
			contains: "DiscountProposalEstimationCreateRequest.ProductID",
			message:  "Field validation for product_id required",
		},
		{
			contains: "DiscountProposalEstimationCreateRequest.ProductQuantity",
			message:  "Field validation for product_quantity required, gt = 0",
		},
		{
			contains: "DiscountProposalEstimationCreateRequest.QuantityType",
			message:  "Field validation for quantity_type required",
		},

		// DISCOUNT PROPOSAL EVENT
		{
			contains: "DiscountProposalEventCreateManyRequest.EventClassID",
			message:  "Field validation for event_class_id required",
		},
		{
			contains: "DiscountProposalEventCreateManyRequest.EventClassName",
			message:  "Field validation for event_class_name required",
		},

		// DISCOUNT PROPOSAL RECIPIENT
		{
			contains: "DiscountProposalRecipientCreateRequest.DiscountProposalID",
			message:  "Field validation for discount_proposal_id required",
		},
		{
			contains: "DiscountProposalRecipientCreateRequest.AccountID",
			message:  "Field validation for account_id required",
		},
		{
			contains: "DiscountProposalRecipientCreateRequest.SubmissionAmount",
			message:  "Field validation for submission_amount required",
		},
	}

	for _, duplicateError := range errors {
		if strings.Contains(err.Error(), duplicateError.contains) {
			message += duplicateError.message + " | "
		}
	}
	if message != "Error : " {
		message += "XXX"
		message = strings.ReplaceAll(message, "| XXX", "")
	} else {
		message += err.Error()
	}

	return message
}
