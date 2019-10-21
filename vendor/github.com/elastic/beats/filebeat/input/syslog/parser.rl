// Code generated by ragel DO NOT EDIT.
package syslog

%%{
    machine syslog;
    write data;
    variable p p;
    variable pe pe;
}%%

var (
  noDuplicates = []byte{'-', '.'}
)

// Parse parses Syslog events.
func Parse(data []byte, event *event) {
    var p, cs int
    pe := len(data)
    tok := 0
    eof := len(data)
    %%{
      action tok {
        tok = p
      }

      action priority {
        event.SetPriority(data[tok:p])
      }

      action message {
        event.SetMessage(data[tok:p])
      }

      action month {
        event.SetMonth(data[tok:p])
      }

      action year{
        event.SetYear(data[tok:p])
      }

      action month_numeric {
        event.SetMonthNumeric(data[tok:p])
      }

      action day {
        event.SetDay(data[tok:p])
      }

      action hour {
        event.SetHour(data[tok:p])
      }

      action minute {
        event.SetMinute(data[tok:p])
      }

      action second {
        event.SetSecond(data[tok:p])
      }

      action nanosecond{
        event.SetNanosecond(data[tok:p])
      }

      # NOTES: This allow to bail out of obvious non valid
      # hostname, this might not be ideal in all situation, but
      # when this happen we just go to the catch all case and at least
      # extract the message
      action lookahead_duplicates{
        if p-1 > 0 {
          for _, b := range noDuplicates {
            if data[p] == b && data[p-1] == b {
              p = tok -1
              fgoto catch_all;
            }
          }
        }
      }

      action hostname {
        event.SetHostname(data[tok:p])
      }

      action program {
        event.SetProgram(data[tok:p])
      }

      action pid {
        event.SetPid(data[tok:p])
      }

      action timezone {
        event.SetTimeZone(data[tok:p])
      }

      action sequence {
        event.SetSequence(data[tok:p])
      }

      include syslog_rfc3164 "syslog_rfc3164.rl";

      write init;
      write exec;
    }%%
}
