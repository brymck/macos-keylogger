#include <CoreFoundation/CoreFoundation.h>
#include <CoreGraphics/CoreGraphics.h>
#include <Carbon/Carbon.h>
#include "keylogger.h"

typedef enum State {
    Up,
    Down,
    Invalid
} State;

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

    //UniChar character = convertKeyCodeToCharacter(keyCode, shift);
    //printf("Character: %lc\n", character);
    TISInputSourceRef currentKeyboard = TISCopyCurrentKeyboardLayoutInputSource();
    CFDataRef layoutData = TISGetInputSourceProperty(currentKeyboard, kTISPropertyUnicodeKeyLayoutData);
    const UCKeyboardLayout *keyboardLayout = (const UCKeyboardLayout *)CFDataGetBytePtr(layoutData);

    UInt32 deadKeyState = 0;
    UniCharCount maxStringLength = 2;
    UniCharCount actualStringLength = 0;
    UniChar unicodeString[maxStringLength];

    // UInt32 modifierKeyState = shiftPressed ? shiftKey : 0;
    OSStatus status = UCKeyTranslate(keyboardLayout, keyCode, kUCKeyActionDown, flags, LMGetKbdType(), 0, &deadKeyState, maxStringLength, &actualStringLength, unicodeString);

    handleButtonEvent((int)keyCode,
                      (int)unicodeString[0],
                      state,
                      ctrl,
                      opt,
                      shift,
                      cmd);

    return event;
}

/*
UniChar convertKeyCodeToCharacter(CGKeyCode keyCode, bool shiftPressed) {
    TISInputSourceRef currentKeyboard = TISCopyCurrentKeyboardLayoutInputSource();
    CFDataRef layoutData = TISGetInputSourceProperty(currentKeyboard, kTISPropertyUnicodeKeyLayoutData);
    const UCKeyboardLayout *keyboardLayout = (const UCKeyboardLayout *)CFDataGetBytePtr(layoutData);

    UInt32 deadKeyState = 0;
    UniCharCount maxStringLength = 255;
    UniCharCount actualStringLength = 0;
    UniChar unicodeString[maxStringLength];

    UInt32 modifierKeyState = shiftPressed ? shiftKey : 0;
    OSStatus status = UCKeyTranslate(keyboardLayout, keyCode, kUCKeyActionDown, modifierKeyState, LMGetKbdType(), 0, &deadKeyState, maxStringLength, &actualStringLength, unicodeString);

    if (status == noErr && actualStringLength > 0) {
        return unicodeString[0];
    } else {
        return 0; // No character found or an error occurred
    }
    return 0;
}
*/

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
