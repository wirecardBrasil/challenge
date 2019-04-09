USE wirecardPayment;

/*Initial inserts*/
INSERT INTO paymentType(
	id, typeDescription) 
    VALUES
    (1, "Boleto"),
    (2, "Cartão de crédito");
    
INSERT INTO paymentState(
	id, stateDescription)
    VALUES
    (1, "Pendente"),
    (2, "Pago"),
    (3, "Cancelado"),
    (4, "Recusado")