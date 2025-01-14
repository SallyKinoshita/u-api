package model

import (
	"testing"
	"time"
)

func TestNewInvoice(t *testing.T) {
	type args struct {
		companyID         int
		businessPartnerID int
		issueDate         time.Time
		dueDate           time.Time
		paymentAmount     float64
	}
	tests := []struct {
		name    string
		args    args
		want    *Invoice
		wantErr bool
	}{
		{
			name: "Valid input",
			args: args{
				companyID:         1,
				businessPartnerID: 2,
				issueDate:         time.Now(),
				dueDate:           time.Now().Add(30 * 24 * time.Hour),
				paymentAmount:     10000,
			},
			want: &Invoice{
				CompanyID:         1,
				BusinessPartnerID: 2,
				Amount:            10000,
				Fee:               400,
				FeeRate:           4.0,
				Tax:               40,
				TaxRate:           10.0,
				TotalAmount:       10440,
				Status:            StatusUnpaid,
			},
			wantErr: false,
		},
		{
			name: "Invalid company ID",
			args: args{
				companyID:         0,
				businessPartnerID: 2,
				issueDate:         time.Now(),
				dueDate:           time.Now().Add(30 * 24 * time.Hour),
				paymentAmount:     10000,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Invalid business partner ID",
			args: args{
				companyID:         1,
				businessPartnerID: 0,
				issueDate:         time.Now(),
				dueDate:           time.Now().Add(30 * 24 * time.Hour),
				paymentAmount:     10000,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Negative payment amount",
			args: args{
				companyID:         1,
				businessPartnerID: 2,
				issueDate:         time.Now(),
				dueDate:           time.Now().Add(30 * 24 * time.Hour),
				paymentAmount:     -10000,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Zero payment amount",
			args: args{
				companyID:         1,
				businessPartnerID: 2,
				issueDate:         time.Now(),
				dueDate:           time.Now().Add(30 * 24 * time.Hour),
				paymentAmount:     0,
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt // 変数のスコープを明示的に指定
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel() // 並列実行を有効化
			got, err := NewInvoice(tt.args.companyID, tt.args.businessPartnerID, tt.args.issueDate, tt.args.dueDate, tt.args.paymentAmount)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewInvoice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && tt.want != nil {
				// 時間を無視して比較
				if got.CompanyID != tt.want.CompanyID ||
					got.BusinessPartnerID != tt.want.BusinessPartnerID ||
					got.Amount != tt.want.Amount ||
					got.Fee != tt.want.Fee ||
					got.Tax != tt.want.Tax ||
					got.TotalAmount != tt.want.TotalAmount ||
					got.FeeRate != tt.want.FeeRate ||
					got.TaxRate != tt.want.TaxRate ||
					got.Status != tt.want.Status {
					t.Errorf("NewInvoice() = %+v, want %+v", got, tt.want)
				}
			}
		})
	}
}
