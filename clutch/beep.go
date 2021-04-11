package clutch

import "time"

/* beep
 */

// /**
// * Function       ModeBEEP
// * @author        Danny
// * @date          2017.07.26
// * @brief         不同模式蜂鸣器不同音调提示
// * @param[in]     void
// * @param[out]    void
// * @retval        void
// * @par History   无
// */
// void ModeBEEP(int mode)
// {
//   int i;
//   pinMode(buzzer, OUTPUT);
//   for (i = 0; i < mode + 1; i++)
//   {
//     digitalWrite(buzzer, LOW); //鸣
//     delay(100);
//     digitalWrite(buzzer, HIGH); //不鸣
//     delay(100);
//   }
//   delay(100);
//   digitalWrite(buzzer, HIGH); //不鸣
// }

// Beep bee time
func Beep() {
	PinBuzzer.Low()
	time.Sleep(1 * time.Second)
	PinBuzzer.High()
}

// Beep bee time
func BeepOn() {
	PinBuzzer.Low()
}

// Beep bee time
func BeepOff() {
	PinBuzzer.High()
}
