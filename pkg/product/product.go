package product

import (
	"fmt"
	"log"
	"regexp"
	"sort"
	"time"
)

const (
	defaultSortField = "CustomerID"
	domain           = "domain"
	hosting          = "hosting"
	email            = "email"
	pdomain          = "pdomain"
	edomain          = "edomain"
)

// Product struct
type Product struct {
	productID       uint64
	CustomerID      string    `json:"customer_id"`
	ProductName     string    `json:"product_name"`
	Domain          string    `json:"domain"`
	StartDate       string    `json:"start_date"`
	DurationMonths  int       `json:"duration_months"`
	emailActiveDate time.Time //Internal only
	emailExpDate    time.Time //Internal only
}

//EmailSchedule response
type EmailSchedule struct {
	CustomerID      string `json:"customer_id"`
	ProductName     string `json:"product_name"`
	Domain          string `json:"domain"`
	EmailActiveDate string `json:"emailActiveDate"`
	EmailExpDate    string `json:"emailExpDate"`
}

//ProductManager - backed by product array
type ProductManager struct {
	Products []Product
}

// Define compare
func (pm *ProductManager) Less(i, j int) bool {
	return pm.Products[i].emailExpDate.After(pm.Products[j].emailExpDate)
}

// Define swap over an array
func (pm *ProductManager) Swap(i, j int) {
	pm.Products[i], pm.Products[j] = pm.Products[j], pm.Products[i]
}

func (pm *ProductManager) Len() int {
	return len(pm.Products)
}

// NewProductManager -Initializing  Product Manager
func NewProductManager() ProductManager {
	// products := []Product{
	// 	Product{uint64(1), "cust4", "domain", "abc.com", "2020-1-1", 18, time.Now(), time.Now()},
	// }
	products := []Product{}
	return ProductManager{
		Products: products,
	}
}

// Process product
func (pm *ProductManager) processProduct(product Product) error {

	switch product.ProductName {
	case domain, pdomain, edomain:
		pm.processDomain(product)
	case hosting:
		pm.processHosting(product)
	case email:
		pm.processemail(product)
	default:
		return fmt.Errorf("Invalid Domain: %s", product.Domain)
	}
	return nil
}

//Billing customer
func (pm *ProductManager) billingCustomer(product Product) error {
	log.Println("Billing the customer")
	return nil
}

// Process domain
func (pm *ProductManager) processDomain(product Product) error {
	pm.billingCustomer(product)
	log.Println("register the domain")
	if (product.Domain) == pdomain {
		log.Println("Securing the domain")
	}
	return nil
}

// Process hosting
func (pm *ProductManager) processHosting(product Product) error {
	pm.billingCustomer(product)
	log.Println("Provisioning account")
	log.Println("Sending welcome email")
	return nil
}

// Process email
func (pm *ProductManager) processemail(product Product) error {
	pm.billingCustomer(product)
	log.Println("Create email routing")
	return nil
}

// GetEmailNotification dates
func (pm *ProductManager) GetEmailNotification(product Product) (time.Time, time.Time) {
	startDate, derr := time.Parse("2006-1-2", product.StartDate)
	log.Printf("Exp  date:%v ", startDate)
	if derr != nil {
		return time.Time{}, time.Time{}
	}
	if product.DurationMonths < 1 {
		return time.Time{}, time.Time{}
	}

	afterActiveDate := time.Time{}

	//Adding months
	expDate := startDate.AddDate(0, product.DurationMonths, 0)
	switch product.ProductName {
	case domain, pdomain, edomain:
		// Set 2 days before expiration

		expDate = expDate.AddDate(0, 0, -2) // 2 days before
		product.emailExpDate = expDate
		log.Printf("Exp -2 days date:%v ", product.emailExpDate)
	case hosting:
		afterActiveDate = startDate.AddDate(0, 0, 1)
		expDate = expDate.AddDate(0, 0, -3) // 3 days before

	case email:
		expDate = expDate.AddDate(0, 0, -1) // 1 days before

	default:
		return time.Time{}, time.Time{}
	}
	return afterActiveDate, expDate
}

// GetEmailSchedule - List email schedule
func (pm *ProductManager) GetEmailSchedule() ([]EmailSchedule, error) {
	// Sort by date -  Need to be done
	sort.SliceStable(pm.Products, func(i, j int) bool {
		return pm.Products[i].CustomerID < pm.Products[j].CustomerID
	})

	emailSch := []EmailSchedule{}

	for _, p := range pm.Products {
		emailActiveDt := ""
		emailExpDt := ""
		if p.emailActiveDate.IsZero() {
			emailActiveDt = ""
		} else {
			emailActiveDt = p.emailActiveDate.Format("2006-1-2")
			log.Printf("emailActiveDt %v", emailActiveDt)
		}

		if p.emailExpDate.IsZero() {
			emailExpDt = ""
		} else {
			emailExpDt = p.emailExpDate.Format("2006-1-2")
			log.Printf("emailExpDt %v", emailExpDt)
		}

		entry := EmailSchedule{
			p.CustomerID,
			p.ProductName,
			p.Domain,
			emailActiveDt,
			emailExpDt,
		}

		emailSch = append(emailSch, entry)
	}
	//sort.Sort(pm.Products())
	return emailSch, nil
}

// Add product
func (pm *ProductManager) Add(product Product) error {

	err := pm.validateProduct(product)
	if err != nil {
		return err
	}

	// Set the start date as today for Add
	currTime := time.Now()
	product.StartDate = currTime.Format("2006-1-2")
	// Set email notification
	startActiveDt, ExpDt := pm.GetEmailNotification(product)
	newProduct := Product{
		0, // Generate Primary key later
		product.CustomerID,
		product.ProductName,
		product.Domain,
		product.StartDate,
		product.DurationMonths,
		startActiveDt,
		ExpDt,
	}

	if err == nil {
		pm.Products = append(pm.Products, newProduct)
		log.Printf("Added Product to the array %v", newProduct)

		log.Printf("Complete Array %v", pm.Products)
	}

	return nil
}

// Load Product
func (pm *ProductManager) Load(product Product) error {
	err := pm.validateProduct(product)
	if err != nil {
		return err
	}

	// Get email notification dates
	startActiveDt, ExpDt := pm.GetEmailNotification(product)

	// Process billing and registering domain
	perr := pm.processProduct(product)
	log.Printf("S:%v,E:%v", startActiveDt, ExpDt)
	newProduct := Product{
		0, // Generate Primary key later
		product.CustomerID,
		product.ProductName,
		product.Domain,
		product.StartDate,
		product.DurationMonths,
		startActiveDt,
		ExpDt,
	}

	if perr == nil && err == nil {
		pm.Products = append(pm.Products, newProduct)
		log.Printf("Added Product to the array %v", newProduct)

	}
	return nil
}

func (pm *ProductManager) GetAll() ([]Product, error) {
	// Sort by customer id - only Natural Sorting
	sort.SliceStable(pm.Products, func(i, j int) bool {
		return pm.Products[i].CustomerID < pm.Products[j].CustomerID
	})

	for _, p := range pm.Products {
		log.Printf("Values %v, %v, %v", p.CustomerID, p.emailExpDate, p.emailActiveDate)
	}
	return pm.Products, nil
}

func (pm *ProductManager) GetByCustomer() ([]Product, error) {
	return nil, nil
}

// Function to validate the product information
func (pm *ProductManager) validateProduct(product Product) error {
	// Based on the product name do the validation
	log.Printf("Inside validateProduct %v", product.ProductName)
	switch product.ProductName {
	case domain, pdomain, edomain, hosting, email:
		err := pm.validateDomain(product)
		return err
		//	case hosting, email:
		// Domain registration required
	default:
		return fmt.Errorf("Invalid product: %s", product.ProductName)
	}
	return nil
}

// Function to validate domain
func (pm *ProductManager) validateDomain(product Product) error {

	if product.ProductName == domain || product.ProductName == pdomain {
		var validDomain, err = regexp.MatchString(`[.com]|[.org]$`, product.Domain)
		// Check for duplicate host registration
		if err != nil || !validDomain {
			return fmt.Errorf("Invalid Domain: %s Error: %v", product.Domain, err)
		}
	}

	if product.ProductName == edomain {
		var valideDomain, err = regexp.MatchString(`[.edu]$`, product.Domain)
		if err != nil || !valideDomain {
			return fmt.Errorf("Invalid eDomain: %s Error: %v", product.Domain, err)
		}
	}

	if !pm.isDupRegistration(product.Domain, product.CustomerID) {
		log.Printf("validateDomain success")
		return nil
	}
	log.Printf("validateDomain failure")
	return fmt.Errorf("Domain: %s is already registered", product.Domain)
}

// IsDupRegistration check the customer got duplicate domain registration
func (pm *ProductManager) isDupRegistration(domainName string, customerID string) bool {
	for _, p := range pm.Products {
		if p.Domain == domainName {
			log.Printf("p.customerid =%v, p.domainName %v", customerID, domainName)
			return true
		}
	}
	return false
}
