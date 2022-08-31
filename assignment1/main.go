package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"text/tabwriter"
)

type Student struct {
	ID         int
	Name       string
	Address    string
	Profession string
	Motivation string
}

func (s *Student) Print() {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 2, '\t', tabwriter.AlignRight)

	fmt.Fprintln(w, "Nomor Absen \t:", s.ID)
	fmt.Fprintln(w, "Nama \t:", s.Name)
	fmt.Fprintln(w, "Alamat \t:", s.Address)
	fmt.Fprintln(w, "Pekerjaan \t:", s.Profession)
	fmt.Fprintln(w, "Alasan Belajar Go \t:", s.Motivation)
	w.Flush()

}

func main() {
	csv := ReadCSV("./assets/class_structure.csv")
	students := GetStudents(csv)

	var id string

	if len(os.Args) < 2 {
		PrintRecords(students)
		os.Exit(0)
	} else {
		id = os.Args[1]
	}

	studentID, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Silahkan masukkan angka absen siswa.")
		os.Exit(1)
	}

	if studentID > len(students) || studentID <= 0 {
		fmt.Printf("Tidak ada siswa dengan nomor absen %d\n", studentID)
		os.Exit(1)
	}

	studentName := students[studentID-1]
	student := &Student{
		ID:         studentID,
		Name:       studentName,
		Address:    "Indonesia",
		Profession: "Mahasiswa",
		Motivation: "Ingin menjadi Backend Engineer",
	}
	student.Print()

}

func ReadCSV(path string) [][]string {
	file, err := os.Open(path)
	defer file.Close()

	if err != nil {
		return nil

	}
	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+path, err)
	}

	return records
}

func GetStudents(records [][]string) []string {
	students := []string{}
	for i := 0; i <= len(records)-1; i++ {
		if records[i][2] == "GLNG-KS04" {
			students = append(students, records[i][3])
		}
	}
	return students
}

func PrintRecords(records []string) {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 2, '\t', tabwriter.AlignRight)
	fmt.Fprintln(w, "NO\tMahasiswa")
	for i := 0; i < len(records); i++ {
		fmt.Fprintf(w, "%v\t%s\n", i+1, records[i])
	}
	w.Flush()
}
