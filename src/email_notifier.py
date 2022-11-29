from fpdf import FPDF


def send_email():
    pass


def generate_email(variante: str):
    email_pdf = FPDF()
    email_pdf.add_page()
    email_pdf.set_font("Arial", size=15)
    if 'a' in variante:
        email_pdf.cell(200, 10, txt="Neue Prüfungsergebnisse!", ln=1, align='C')
        if 'b' in variante:
            email_pdf.cell(200, 10, txt="Prüfung bestanden!", ln=2, align='C')
        if 'c' in variante:
            email_pdf.cell(200, 10, txt="Durchgefallen!", ln=2, align='C')
        else:
            email_pdf.cell(200, 10, txt="Bitte öffne Campus Dual!", ln=2, align='C')
        email_pdf.output("Email.pdf")
    elif 'd' in variante:
        email_pdf.cell(200, 10, txt="Neue Anmeldungen verfügbar!", ln=1, align='C')
        email_pdf.output("Email.pdf")
