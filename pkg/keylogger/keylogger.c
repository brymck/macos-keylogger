#include <CoreFoundation/CoreFoundation.h>
#include <CoreGraphics/CoreGraphics.h>
#include <Carbon/Carbon.h>
#include "keylogger.h"

typedef enum State {
    Up,
    Down,
    Invalid
} State;

static const UniCharCount MAX_STRING_LENGTH = 4;

extern void handleButtonEvent(int k,
                              int ch,
                              State s,
                              bool ctrl,
                              bool opt,
                              bool shift,
                              bool cmd);

// The following callback method is invoked on every keypress.
static inline CGEventRef CGEventCallback(CGEventTapProxy proxy,
                                         CGEventType type,
                                         CGEventRef event,
                                         void *refcon) {
    // Validate the input event
    State state;
    if (type == kCGEventKeyDown) {
        state = Down;
    } else if (type == kCGEventKeyUp) {
        state = Up;
    } else if (type == kCGEventFlagsChanged) {
        state = Invalid;
    } else {
        return event;
    }

    // Retrieve the key code
    CGKeyCode keyCode = (CGKeyCode) CGEventGetIntegerValueField(event, kCGKeyboardEventKeycode);

    // Detect modifier keys
    const CGEventFlags flags = CGEventGetFlags(event);

    bool ctrl = (flags & kCGEventFlagMaskControl) != 0;
    bool opt = (flags & kCGEventFlagMaskAlternate) != 0;
    bool shift = (flags & kCGEventFlagMaskShift) != 0;
    bool cmd = (flags & kCGEventFlagMaskCommand) != 0;

    // TODO Just use Carbon since I obviously need to do some conversion anyway
    UInt16 modifierKeyState = shift << 1 | ctrl << 2 | opt << 3 | cmd << 4;

    TISInputSourceRef currentKeyboard = TISCopyCurrentKeyboardLayoutInputSource();
    CFDataRef layoutData = TISGetInputSourceProperty(currentKeyboard, kTISPropertyUnicodeKeyLayoutData);
    const UCKeyboardLayout *keyboardLayout = (UCKeyboardLayout *)CFDataGetBytePtr(layoutData);
    static UInt32 deadKeyState = 0;
    UniCharCount actualStringLength = 0;
    UniChar unicodeString[MAX_STRING_LENGTH];

    OSStatus status = UCKeyTranslate(keyboardLayout,
                                     keyCode,
                                     kUCKeyActionDisplay,
                                     modifierKeyState,
                                     LMGetKbdType(),
                                     kUCKeyTranslateNoDeadKeysBit,
                                     &deadKeyState,
                                     MAX_STRING_LENGTH,
                                     &actualStringLength,
                                     unicodeString);

    handleButtonEvent((int)keyCode,
                      (int)unicodeString[0],
                      state,
                      ctrl,
                      opt,
                      shift,
                      cmd);

    return event;
}

static inline void listen() {
    CGEventMask eventMask = CGEventMaskBit(kCGEventKeyDown) | CGEventMaskBit(kCGEventKeyUp);


    CFMachPortRef eventTap = CGEventTapCreate(kCGSessionEventTap,
                                              kCGHeadInsertEventTap,
                                              0,
                                              eventMask,
                                              CGEventCallback,
                                              NULL);

    if (!eventTap) {
        fprintf(stderr, "ERROR: Unable to create event tap.");
        exit(1);
    }

    CFRunLoopSourceRef runLoopSource = CFMachPortCreateRunLoopSource(kCFAllocatorDefault, eventTap, 0);
    CFRunLoopAddSource(CFRunLoopGetCurrent(),
                       runLoopSource,
                       kCFRunLoopCommonModes);
    CGEventTapEnable(eventTap, true);

    CFRunLoopRun();
}
