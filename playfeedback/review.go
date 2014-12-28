package playfeedback

import (
    "os"
    "bufio"
    "encoding/csv"
    "fmt"
    "strconv"
    "time"
)

type Review struct {
    Link    string
    Text    string
    Title   string
    Rating  int
    Created time.Time
    Updated *time.Time
    Version int
    Device  string
}

func FromCsvFile(filename string) (rx []Review, err error) {
    file, err := os.Open(filename)
    
    if err != nil {
        fmt.Printf("Error opening file %s : %v\n", filename, err)
        return nil, err
    }
    
    // map names to indeces
    n2i := make(map[string]int)
    bf := bufio.NewReader(file)
    r  := csv.NewReader(bf)
    r.LazyQuotes = true
    r.TrimLeadingSpace = true
    
    records, err := r.ReadAll()
    if err != nil {
        fmt.Printf("Error reading CSV: %v\n", err)
        return nil, err
    }
    
    // Print available columns
    for i,name := range records[0] {
        n2i[name] = i
        fmt.Printf("%d: %s\n", i, name)
    }
    
    rx = make([]Review, len(records) - 1)
    for i := 1; i < len(records); i++ {
        r := records[i]
        rating,_   := strconv.Atoi(r[n2i["Star Rating"]])
        
        createdstr := r[n2i["Review Submit Date and Time"]]
        created,_  := time.Parse(time.RFC3339, createdstr)
        
        var updated *time.Time
        updatedstr := r[n2i["Review Last Update Date and Time"]]
        if (len(updatedstr) > 0) {
            t,_ := time.Parse(time.RFC3339, updatedstr)
            updated = &t
        }
        version,_ := strconv.Atoi(r[n2i["App Version Code"]])
        device    := r[n2i["Device"]]
        
        rx[i-1] = Review {
            r[n2i["Review Link"]],
            r[n2i["Review Text"]],
            r[n2i["Review Title"]],
            rating,
            created,
            updated,
            version,
            device,
        }
    }   
    fmt.Printf("Read %d reviews\n", len(rx))
     
    return rx, nil
}

func FilterRecent(rx []Review) []Review {
    recent := make([]Review,0)
    for _,r := range rx {
        var t time.Time
        if r.Updated != nil {
            t = *r.Updated
        } else {
            t = r.Created
        }
        since := time.Since(t)
        if since < 24*time.Hour {
            recent = append(recent, r)
        }
    }
    return recent
}