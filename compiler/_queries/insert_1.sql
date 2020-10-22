/*
  @param stock -> ((symbol, name, market, currency, enabled)...)
  @exec
*/
-- AddStocks adds provided stocks to the database
insert into stocks (symbol, name, market, currency) values :stock;
