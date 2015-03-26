#!/usr/bin/env ruby

this_dir = File.expand_path(File.dirname(__FILE__))
lib_dir = File.dirname(this_dir)
$LOAD_PATH.unshift(lib_dir) unless $LOAD_PATH.include?(lib_dir)

require 'grpc'
require 'xrates_services'

include Xrates

def main
  stub = XRates::Stub.new('localhost:50800')

  puts "Getting one currency"
  rate = stub.get(Currency.new(currency: 'MYR'))
  puts "Rate for MYR: #{rate.rate}"

  puts "Getting multiple currencies"
  rates = stub.all(Currencies.new(currencies: []))
  rates.rates.each do |cur, rate|
    puts "Rate for #{cur}: #{rate}"
  end
end

main
