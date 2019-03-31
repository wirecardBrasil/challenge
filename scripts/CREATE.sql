/*DROP DATABASE wirecardPayment;*/

CREATE DATABASE IF NOT EXISTS wirecardPayment;
USE  wirecardPayment;

/*To storage client's data */
CREATE TABLE IF NOT EXISTS regClient(
	id BIGINT AUTO_INCREMENT NOT NULL,
    CONSTRAINT PK_regClient_id PRIMARY KEY(id),
    clientName VARCHAR(100) NOT NULL,
    email VARCHAR(50),
    cpfCnpj VARCHAR(14) UNIQUE
);

CREATE TABLE IF NOT EXISTS regBuyer(
	id BIGINT AUTO_INCREMENT NOT NULL,
    CONSTRAINT PK_regBuyert_id PRIMARY KEY(id),
    clientName VARCHAR(100) NOT NULL,
    email VARCHAR(50) NOT NULL,
    cpfCnpj VARCHAR(14) NOT NULL UNIQUE
);

/*Index to otimize queries by cpf/cnpj*/
CREATE INDEX regClient_cpfCnpj ON regClient(cpfCnpj);
CREATE INDEX regBuyer_cpfCnpj ON regBuyer(cpfCnpj);

/*To storage payment types*/
CREATE TABLE IF NOT EXISTS paymentType(
	id SMALLINT NOT NULL,
    CONSTRAINT PK_paymentType_id PRIMARY KEY(id),
    typeDescription VARCHAR(50) NOT NULL
);

/*To store payment's states */
CREATE TABLE IF NOT EXISTS paymentState(
	id SMALLINT NOT NULL,
    CONSTRAINT PK_paymentState_id PRIMARY KEY(id),
    stateDescription VARCHAR(50) NOT NULL
);

/*To storage payment*/
CREATE TABLE IF NOT EXISTS payment(
	id BIGINT AUTO_INCREMENT NOT NULL,
    CONSTRAINT PK_payment_id PRIMARY KEY(id),
    idClient BIGINT NOT NULL,
    CONSTRAINT FK_payment_idClient FOREIGN KEY(idClient)
		REFERENCES regClient(id),
	idBuyer BIGINT NOT NULL,
    CONSTRAINT FK_payment_idBuyer FOREIGN KEY(idBuyer)
		REFERENCES regBuyer(id),
	idPaymentType SMALLINT NOT NULL,
	CONSTRAINT FK_payment_idPaymentType FOREIGN KEY(idPaymentType)
		REFERENCES paymentType(id),
	amount DECIMAL(15, 2) NOT NULL
);

/*To store card infos */
CREATE TABLE IF NOT EXISTS cardPayment(
	id BIGINT AUTO_INCREMENT NOT NULL,
    CONSTRAINT PK_cardPayment_id PRIMARY KEY(id),
    idPayment BIGINT NOT NULL,
    CONSTRAINT FK_cardPayment_idPayment FOREIGN KEY(idPayment)
		REFERENCES payment(id),
	holderName VARCHAR(100) NOT NULL,
    cardFinalNumber VARCHAR(4) NOT NULL, /*store only the final number for safety issues*/
    expirationDate DATE NOT NULL
    /*cvv - Check if it's necessary to store, because it could be unsafe*/
);

/*To store boleto's infos*/
CREATE TABLE IF NOT EXISTS boletoPayment(
	id BIGINT AUTO_INCREMENT NOT NULL,
    CONSTRAINT PK_boletoPayment_id PRIMARY KEY(id),
    idPayment BIGINT NOT NULL,
    CONSTRAINT FK_boletoPayment_idPayment FOREIGN KEY(idPayment)
		REFERENCES payment(id),
	boletoNumber VARCHAR(50) NOT NULL /*current size is 48, verify*/
)



